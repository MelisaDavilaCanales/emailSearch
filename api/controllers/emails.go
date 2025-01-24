package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/MelisaDavilaCanales/emailSearch/api/constant"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/MelisaDavilaCanales/emailSearch/api/storage"
)

var (
	GetEmailsFunc = storage.GetEmails
	GetEmailFunc  = storage.GetMail

	notFoundErr        *models.NotFoundError
	mssgEmailNotFound  = "Email not found"
	mssgEmailsNotFound = "Emails not found"
	mssSearchSuccess   = "Search successfully"
)

func InitMockGetEmails(emailHistData *models.EmailHitsData, err error) {
	GetEmailsFunc = func(_ models.SearchParams) (*models.EmailHitsData, error) {
		return emailHistData, err
	}
}

func DisableMockGetEmails() {
	GetEmailsFunc = storage.GetEmails
}

func GetEmails(w http.ResponseWriter, r *http.Request) {
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

	emailHitsData, err := GetEmailsFunc(searchParams)
	if err != nil {
		if errors.As(err, &notFoundErr) {
			data := models.NewEmailsResponseData(0, 0, 0, nil)
			response := models.NewResponse(mssgEmailsNotFound, data)
			render.JSON(w, r, response)

			return
		}

		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting emails", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)

		return
	}

	if emailHitsData == nil || emailHitsData.Total.Value == 0 {
		data := models.NewEmailsResponseData(0, 0, 0, nil)
		response := models.NewResponse(mssgEmailsNotFound, data)
		render.JSON(w, r, response)

		return
	}

	totalPages := getTotalPages(emailHitsData.Total.Value, pagination.MaxResults)
	if pagination.PageNumber > totalPages {
		data := models.NewEmailsResponseData(totalPages, pagination.PageNumber, pagination.PageSize, nil)
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

	data := models.NewEmailsResponseData(totalPages, pagination.PageNumber, pagination.PageSize, emails)
	response := models.NewResponse(mssSearchSuccess, data)
	render.JSON(w, r, response)
}

func InitMockGetEmail(email *models.Email, err error) {
	GetEmailFunc = func(_ string) (*models.Email, error) {
		return email, err
	}
}

func DisableMockGetEmail() {
	GetEmailFunc = storage.GetMail
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, constant.ID_PARAM)

	email, err := storage.GetMail(id)
	if err != nil {
		if errors.As(err, &notFoundErr) {
			response := models.NewResponse(mssgEmailNotFound, nil)
			render.JSON(w, r, response)

			return
		}

		responseError := models.NewResponseError(http.StatusInternalServerError, "Error getting email", err)
		http.Error(w, responseError.Error(), responseError.StatusCode)
		fmt.Println(err)

		return
	}

	if email == nil {
		response := models.NewResponse(mssgEmailNotFound, nil)
		render.JSON(w, r, response)

		return
	}

	newEmail := models.EmailResponse{
		ID:        email.ID,
		MessageID: email.MessageID,
		Date:      email.Date,
		From:      email.From,
		To:        email.To,
		Subject:   email.Subject,
		Cc:        email.Cc,
		XFolder:   email.XFolder,
		XOrigin:   email.XOrigin,
		XFileName: email.XFileName,
		Content:   email.Content,
	}

	response := models.NewResponse(mssSearchSuccess, newEmail)
	render.JSON(w, r, response)
}
