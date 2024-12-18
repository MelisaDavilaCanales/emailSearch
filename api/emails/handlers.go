package emails

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/MelisaDavilaCanales/emailSearch/api/storage"
	"github.com/MelisaDavilaCanales/emailSearch/api/utils"
)

func GetEmails(w http.ResponseWriter, r *http.Request) {
	pageNumberStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	searchTerm := r.URL.Query().Get("term")
	searchfield := r.URL.Query().Get("field")

	pageNumber, pageSize, resultsFrom, maxResults := utils.ProcessPaginatedParams(pageNumberStr, pageSizeStr)

	emailHitsData, err := storage.GetEmails(searchTerm, searchfield, resultsFrom, maxResults)
	if err != nil {
		responseError := models.NewResponseError(http.StatusInternalServerError, "Error searching emails", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)

		return
	}

	if emailHitsData == nil || emailHitsData.Total.Value == 0 {
		data := models.NewEmailsResponseData(0, 0, 0, []models.EmailSummary{})
		response := models.NewResponse("No emails found", data)
		render.JSON(w, r, response)

		return
	}

	totalPages := utils.GetTotalPages(emailHitsData.Total.Value, maxResults)
	if pageNumber > totalPages {
		data := models.NewEmailsResponseData(totalPages, pageNumber, pageSize, []models.EmailSummary{})
		response := models.NewResponse("Page out of range", data)
		render.JSON(w, r, response)

		return
	}

	emails := make([]models.EmailSummary, len(emailHitsData.Hits))
	for i, hit := range emailHitsData.Hits {
		emails[i] = models.EmailSummary{
			Id:      hit.ID,
			Date:    hit.Email.Date,
			From:    hit.Email.From,
			To:      hit.Email.To,
			Subject: hit.Email.Subject,
		}
	}

	data := models.NewEmailsResponseData(totalPages, pageNumber, pageSize, emails)
	response := models.NewResponse("Emails found successfully", data)
	render.JSON(w, r, response)
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	email, err := storage.GetMail(id)
	if err != nil {
		if err.StatusCode == http.StatusNotFound {
			response := models.NewResponse("Email not found", nil)
			render.JSON(w, r, response)

			return
		}

		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting email", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)

		return
	}

	response := models.NewResponse("Email found successfully", email)
	render.JSON(w, r, response)
}
