package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/emails"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/models"
	models_wp "github.com/MelisaDavilaCanales/emailSearch/indexer/models/workerpool"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/persons"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/profiling"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/storage"
)

var CantGoRutine int

func main() {

	cpuFile, memFile, traceFile := profiling.StartProfiling()
	defer profiling.StopProfiling(cpuFile, memFile, traceFile)

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Number of CPUs: ", runtime.NumCPU())

	timeInit := time.Now()

	directory, err := emails.GetDirectory()
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
		defer wgProcessDir.Done()

		var wgProcessEmailDirectory sync.WaitGroup

		files, err := os.ReadDir(directory)
		if err != nil {
			fmt.Printf("Error reading top-level directory: %v\n", err)
			return
		}

		for _, file := range files {
			if file.IsDir() {
				subDirPath := filepath.Join(directory, file.Name())

				wgProcessEmailDirectory.Add(1)
				CantGoRutine++
				go func(path string) {
					defer wgProcessEmailDirectory.Done()
					err := emails.ProcessSubDirectory(path, emailPathCh)
					if err != nil {
						fmt.Printf("Error processing subdirectory %s: %v\n", path, err)
					}
				}(subDirPath)
			}
		}

		wgProcessEmailDirectory.Wait()
		close(emailPathCh)
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
	fmt.Printf("--- GoRutines ---: %v\n", CantGoRutine)
	fmt.Printf("--- EMAIL_BATCH_SIZE ---: %v\n", constant.EMAIL_BATCH_SIZE)
	fmt.Printf("--- PROCESS_EMAILS_WORKERS_COUNT ---: %v\n", constant.PROCESS_EMAILS_WORKERS_COUNT)
	fmt.Printf("--- SEND_EMAILS_WORKERS_COUNT ---: %v\n", constant.SEND_EMAILS_WORKERS_COUNT)
	fmt.Printf("--- STRUCTURE_PERSONS_WORKERS_COUNT ---: %v\n", constant.STRUCTURE_PERSONS_WORKERS_COUNT)
	fmt.Printf("--- BUFFER_CAPACITY ---: %v\n", constant.BUFFER_CAPACITY)
}
