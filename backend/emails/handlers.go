package emails

import (
	"backend/models"
	"backend/storage"
	"backend/utils"
	"encoding/json"
	"math"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetEmail(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	email, err := storage.GetMail(id)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting email", err)

		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	data := map[string]interface{}{
		"email": email,
	}

	response := models.NewResponse(http.StatusOK, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func GetAllEmails(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, pageSize, from, max := utils.ProcessPaginatedParams(pageStr, pageSizeStr)

	hitsData, err := storage.GetAllEmails(from, max)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting emails", err)

		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	totalPages := int(math.Ceil(float64(hitsData.Total.Value) / float64(max)))
	data := map[string]interface{}{
		"totalPages": totalPages,
		"page":       page,
		"pageSize":   pageSize,
		"emails":     hitsData.Hits,
	}

	response := models.NewResponse(http.StatusOK, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	term := r.URL.Query().Get("term")
	field := r.URL.Query().Get("field")

	page, pageSize, from, max := utils.ProcessPaginatedParams(pageStr, pageSizeStr)

	hitsData, err := storage.SearchEmails(term, field, from, max)
	if err != nil {
		// responseError := models.NewResponseError(http.StatusInternalServerError, "Error searching emails", err)

		http.Error(w, "responseError.Error()", 700)
	}

	totalPages := int(math.Ceil(float64(hitsData.Total.Value) / float64(max)))
	data := map[string]interface{}{
		"totalPages": totalPages,
		"page":       page,
		"pageSize":   pageSize,
		"emails":     hitsData.Hits,
	}

	response := models.NewResponse(http.StatusOK, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
