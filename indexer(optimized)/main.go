package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/emails"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/models"
	models_wp "github.com/MelisaDavilaCanales/emailSearch/indexer/models/workerpool"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/persons"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/storage"
)

func main() {
	// profiling.StartHTTPProfiler()

	// cpuFile, memFile := profiling.StartProfiling()
	// defer profiling.StopProfiling(cpuFile, memFile)

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Number of CPUs: ", runtime.NumCPU())

	timeInit := time.Now()

	_, err := emails.GetDirectory()
	if err != nil {
		log.Fatal("Error getting directory: ", err)
	}

	if err := config.SetEnvVars(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	if err := storage.TryConnectionAPI(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := storage.CreateEmailIndex(); err != nil {
		log.Fatal("Error creating email index: ", err)
	}

	if err := storage.CreatePersonIndex(); err != nil {
		log.Fatal("Error creating person index: ", err)
	}

	emailPathCh := make(chan string, constant.BUFFER_CAPACITY)
	emailStructCh1 := make(chan models_wp.Result[*models.EmailData], constant.BUFFER_CAPACITY)
	emailStructCh2 := make(chan models_wp.Result[*models.EmailData], constant.BUFFER_CAPACITY)

	var wgProcessDir, wgProcessEmailFiles, wgProcessAndSendEmails, wgStructurePersons, wgBuildAndSendPersonBulk sync.WaitGroup

	wgProcessDir.Add(1)

	go func() {
		defer close(emailPathCh)
		defer wgProcessDir.Done()

		err := emails.ProcessEmailDirectory(emailPathCh)
		if err != nil {
			log.Fatal("Error processing email directory: ", err, " Please provide a valid directory, example: go run main.go <directory>")
		}
	}()

	wpProcessEmailsFiles := models_wp.NewWorkerPool(emails.ProcessEmailsFiles, &wgProcessEmailFiles, constant.PROCESS_EMAILS_WORKERS_COUNT, emailPathCh, emailStructCh1, emailStructCh2)
	wpProcessEmailsFiles.Start()

	wpProcessAndSendEmails := models_wp.NewWorkerPool(emails.ProcessAndSendEmails, &wgProcessAndSendEmails, constant.SEND_EMAILS_WORKERS_COUNT, emailStructCh1)
	wpProcessAndSendEmails.Start()

	wpStructurePersons := models_wp.NewWorkerPool(persons.StructurePersons, &wgStructurePersons, constant.STRUCTURE_PERSONS_WORKERS_COUNT, emailStructCh2)
	wpStructurePersons.Start()

	wgBuildAndSendPersonBulk.Add(1)

	go func() {
		wgStructurePersons.Wait()

		defer wgBuildAndSendPersonBulk.Done()
		persons.BuildAndSendPersonBulk()
	}()

	wgProcessEmailFiles.Wait()
	wgProcessAndSendEmails.Wait()
	wgBuildAndSendPersonBulk.Wait()

	timeSince := time.Since(timeInit)
	fmt.Println("Indexing Time: ", timeSince)
	fmt.Printf("TotalEmails: %v\n", emails.TotalEmails)
	fmt.Printf("TotalEmailsValid: %v\n", emails.TotalEmailsValid)
	fmt.Printf("TotalEmailsInvalid: %v\n", emails.TotalEmailsInvalid)
}
