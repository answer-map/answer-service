package app

import (
	"encoding/json"
	"net/http"

	"github.com/answer-map/answer-service/service"
)

const (
	correlationIDHeaderName string = "X-Correlation-ID"
	contentTypeHeaderName   string = "Content-Type"
	contentTypeHeaderJSON   string = "application/json"
)

func render(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func renderMarshalJSON(w http.ResponseWriter, statusCode int, body any) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Add(contentTypeHeaderName, contentTypeHeaderJSON)
	w.Write(bodyBytes)
}

func renderServiceError(w http.ResponseWriter, err error) {
	switch err {
	case service.ErrBadRequest:
		render(w, http.StatusBadRequest)
	case service.ErrResourceNotFound:
		render(w, http.StatusNotFound)
	case service.ErrConflict:
		render(w, http.StatusConflict)
	default:
		render(w, http.StatusInternalServerError)
	}
}
