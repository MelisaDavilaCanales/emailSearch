package persons

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"backend/models"
	"backend/storage"
	"backend/utils"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	pageNumberStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	searchTerm := r.URL.Query().Get("term")
	searchfield := r.URL.Query().Get("field")

	page, pageSize, resultsFrom, maxResults := utils.ProcessPaginatedParams(pageNumberStr, pageSizeStr)

	personHitsData, err := storage.GetPersons(searchTerm, searchfield, resultsFrom, maxResults)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error searching persons", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	if personHitsData == nil {
		return
	}

	if personHitsData.Total.Value == 0 {
		data := models.NewPersonResponseData(0, 0, 0, []models.Person{})
		response := models.NewResponse("No persons found", data)
		render.JSON(w, r, response)

		return
	}

	totalPages := utils.GetTotalPages(personHitsData.Total.Value, maxResults)
	if page > totalPages {
		data := models.NewEmailsResponseData(totalPages, page, pageSize, []models.EmailSummary{})
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
	response := models.NewResponse("Persons found successfully", data)
	render.JSON(w, r, response)
}
