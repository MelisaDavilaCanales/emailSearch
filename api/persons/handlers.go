package persons

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/MelisaDavilaCanales/emailSearch/api/constant"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/MelisaDavilaCanales/emailSearch/api/storage"
	"github.com/MelisaDavilaCanales/emailSearch/api/utils"
)

var (
	notFound            *models.NotFoundError
	mssgPersonsNotFound = "Persons not found"
	mssSearchSuccess    = "Search successfully"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	pageNumberStr := r.URL.Query().Get(constant.PAGE_NUMBER_PARAM)
	pageSizeStr := r.URL.Query().Get(constant.PAGE_SIZE_PARAM)
	searchTerm := r.URL.Query().Get(constant.SEARCH_TERM_PARAM)
	searchfield := r.URL.Query().Get(constant.SEARCH_FIELD_PARAM)

	page, pageSize, resultsFrom, maxResults := utils.ProcessPaginatedParams(pageNumberStr, pageSizeStr)

	personHitsData, err := storage.GetPersons(searchTerm, searchfield, resultsFrom, maxResults)
	if err != nil {
		if errors.As(err, &notFound) {
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

	totalPages := utils.GetTotalPages(personHitsData.Total.Value, maxResults)
	if page > totalPages {
		data := models.NewPersonResponseData(totalPages, page, pageSize, nil)
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
		fmt.Println(persons[i].Name)
	}

	data := models.NewPersonResponseData(totalPages, page, pageSize, persons)
	response := models.NewResponse(mssSearchSuccess, data)
	render.JSON(w, r, response)
}
