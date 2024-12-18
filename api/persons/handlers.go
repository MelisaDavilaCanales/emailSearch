package persons

import (
	"backend/models"
	"backend/storage"
	"backend/utils"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func GetPersonsHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	term := r.URL.Query().Get("term")
	field := r.URL.Query().Get("field")

	page, pageSize, from, max := utils.ProcessPaginatedParams(pageStr, pageSizeStr)

	personHitsData, err := storage.GetPersons(term, field, from, max)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error searching persons", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	if personHitsData == nil {
		return
	}

	if personHitsData.Total.Value == 0 {
		data := models.NewPersonResponseData(0, 0, 0, []models.Person{})
		response := models.Response{
			Message: "No persons found",
			Data:    data,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	totalPages := int(math.Ceil(float64(personHitsData.Total.Value) / float64(max)))
	if page > totalPages {
		data := models.NewEmailsResponseData(totalPages, page, pageSize, []models.EmailSummary{})
		response := models.NewResponse("Page out of range", data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

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
	response := models.Response{
		Message: "Persons found successfully",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
