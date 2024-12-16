package persons

import (
	"backend/models"
	"backend/storage"
	"backend/utils"
	"encoding/json"
	"math"
	"net/http"
)

func GetAllPersons(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, pageSize, from, max := utils.ProcessPaginatedParams(pageStr, pageSizeStr)

	hitsData, err := storage.GetAllPersons(from, max)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting persons", err)

		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	totalPages := int(math.Ceil(float64(hitsData.Total.Value) / float64(max)))
	data := map[string]interface{}{
		"totalPages": totalPages,
		"page":       page,
		"pageSize":   pageSize,
		"persons":    hitsData.Hits,
	}

	response := models.NewResponse(http.StatusOK, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)

}

func SearchPersons(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	term := r.URL.Query().Get("term")
	field := r.URL.Query().Get("field")

	page, pageSize, from, max := utils.ProcessPaginatedParams(pageStr, pageSizeStr)

	hitsData, err := storage.SearchPersons(term, field, from, max)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error searching persons", err)

		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	totalPages := int(math.Ceil(float64(hitsData.Total.Value) / float64(max)))
	data := map[string]interface{}{
		"totalPages": totalPages,
		"page":       page,
		"pageSize":   pageSize,
		"person":     hitsData.Hits,
	}

	response := models.NewResponse(http.StatusOK, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
