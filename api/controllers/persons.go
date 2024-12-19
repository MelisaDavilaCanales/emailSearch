package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/MelisaDavilaCanales/emailSearch/api/storage"
)

var (
	notFoundErr         *models.NotFoundError
	mssgPersonsNotFound = "Persons not found"
	mssgSearchSuccess   = "Search successfully"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	queryParams := getQueryParams(r)
	pagination := processPaginatedParams(queryParams.PageNumber, queryParams.PageSize)

	searchParams := models.SearchParams{
		SearchTerm:  queryParams.SearchTerm,
		SearchField: queryParams.SearchField,
		SortField:   queryParams.SortField,
		SortOrder:   queryParams.SortOrder,
		ResultsFrom: pagination.ResultsFrom,
		MaxResults:  pagination.PageSize,
	}

	personHitsData, err := storage.GetPersons(searchParams)
	if err != nil {
		if errors.As(err, &notFoundErr) {
			data := models.NewPersonResponseData(0, 0, 0, nil)
			response := models.NewResponse(mssgPersonsNotFound, data)
			render.JSON(w, r, response)

			return
		}
	}

	if personHitsData == nil || personHitsData.Total.Value == 0 {
		data := models.NewPersonResponseData(0, 0, 0, nil)
		response := models.NewResponse(mssgPersonsNotFound, data)
		render.JSON(w, r, response)

		return
	}

	totalPages := getTotalPages(personHitsData.Total.Value, pagination.MaxResults)
	if pagination.PageNumber > totalPages {
		data := models.NewPersonResponseData(totalPages, pagination.PageNumber, pagination.PageSize, nil)
		response := models.NewResponse("Page out of range", data)
		render.JSON(w, r, response)

		return
	}

	persons := make([]models.Person, len(personHitsData.Hits))
	for i, personHit := range personHitsData.Hits {
		persons[i] = models.Person{
			Id:    personHit.ID,
			Name:  personHit.Person.Name,
			Email: personHit.Person.Email,
		}
	}

	data := models.NewPersonResponseData(totalPages, pagination.PageNumber, pagination.PageSize, persons)
	response := models.NewResponse(mssgSearchSuccess, data)
	render.JSON(w, r, response)
}