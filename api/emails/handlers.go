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

func GetEmails(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	term := r.URL.Query().Get("term")
	field := r.URL.Query().Get("field")

	page, pageSize, from, max := utils.ProcessPaginatedParams(pageStr, pageSizeStr)

	emailHitsData, err := storage.GetEmails(term, field, from, max)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error searching emails", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	if emailHitsData.Total.Value == 0 {
		data := models.NewEmailsResponseData(0, 0, 0, []models.EmailSummary{})
		response := models.Response{
			Message: "No emails found",
			Data:    data,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	totalPages := int(math.Ceil(float64(emailHitsData.Total.Value) / float64(max)))

	emails := make([]models.EmailSummary, 0)
	for _, hit := range emailHitsData.Hits {
		emails = append(emails, models.EmailSummary{
			Id:      hit.ID,
			Date:    hit.Email.Date,
			From:    hit.Email.From,
			To:      hit.Email.To,
			Subject: hit.Email.Subject,
		})
	}

	data := models.NewEmailsResponseData(totalPages, page, pageSize, emails)
	response := models.Response{
		Message: "Emails found successfully",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetEmail(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	email, err := storage.GetMail(id)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting email", err)

		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(email)
}
