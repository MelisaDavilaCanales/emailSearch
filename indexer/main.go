package main

import (
	"fmt"
	"log"
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
	timeInit := time.Now()

	if err := config.SetEnvVars(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	if err := storage.CreateEmailIndex(); err != nil {
		log.Fatal("Error creating email index: ", err)
	}

	if err := storage.CreatePersonIndex(); err != nil {
		log.Fatal("Error creating person index: ", err)
	}

	emailPathCh := make(chan string, constant.BUFFER_CAPACITY)
	emailStructCh1 := make(chan models_wp.Result[*models.Email], constant.BUFFER_CAPACITY)
	emailStructCh2 := make(chan models_wp.Result[*models.Email], constant.BUFFER_CAPACITY)

	var wgProcessDir, wgProcessEmailFiles, wgProcessAndSendEmails, wgStructurePersons, wgBuildAndSendPersonBulk sync.WaitGroup

	wgProcessDir.Add(1)

	go func() {
		defer close(emailPathCh)
		defer wgProcessDir.Done()
		emails.ProcessEmailDirectory(emailPathCh)
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
}
