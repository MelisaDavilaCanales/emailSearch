package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/MelisaDavilaCanales/emailSearch/api/constant"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/MelisaDavilaCanales/emailSearch/api/storage"
)

var (
	notFound           *models.NotFoundError
	mssgEmailNotFound  = "Email not found"
	mssgEmailsNotFound = "Emails not found"
	mssSearchSuccess   = "Search successfully"
)

func GetEmails(w http.ResponseWriter, r *http.Request) {
	pageNumberStr := r.URL.Query().Get(constant.PAGE_NUMBER_PARAM)
	pageSizeStr := r.URL.Query().Get(constant.PAGE_SIZE_PARAM)
	searchTerm := r.URL.Query().Get(constant.SEARCH_TERM_PARAM)
	searchfield := r.URL.Query().Get(constant.SEARCH_FIELD_PARAM)

	pageNumber, pageSize, resultsFrom, maxResults := ProcessPaginatedParams(pageNumberStr, pageSizeStr)

	emailHitsData, err := storage.GetEmails(searchTerm, searchfield, resultsFrom, maxResults)
	if err != nil {
		if errors.As(err, &notFound) {
			data := models.NewEmailsResponseData(0, 0, 0, nil)
			response := models.NewResponse(mssgEmailsNotFound, data)
			render.JSON(w, r, response)

			return
		}

		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting emails", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)
	}

	if emailHitsData == nil || emailHitsData.Total.Value == 0 {
		data := models.NewEmailsResponseData(0, 0, 0, nil)
		response := models.NewResponse(mssgEmailsNotFound, data)
		render.JSON(w, r, response)

		return
	}

	totalPages := GetTotalPages(emailHitsData.Total.Value, maxResults)
	if pageNumber > totalPages {
		data := models.NewEmailsResponseData(totalPages, pageNumber, pageSize, nil)
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
	response := models.NewResponse(mssSearchSuccess, data)
	render.JSON(w, r, response)
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, constant.ID_PARAM)

	email, err := storage.GetMail(id)
	if err != nil {
		if errors.As(err, &notFound) {
			response := models.NewResponse(mssgEmailNotFound, nil)
			render.JSON(w, r, response)

			return
		}

		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting email", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)

		return
	}

	if email == nil {
		response := models.NewResponse(mssgEmailNotFound, nil)
		render.JSON(w, r, response)

		return
	}

	response := models.NewResponse(mssSearchSuccess, email)
	render.JSON(w, r, response)
}
