package app

import (
	"encoding/json"
	"net/http"

	"github.com/answer-map/answer-service/model"
	"go.uber.org/zap"
)

func (app *App) createAnswer(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logger := app.logger
	logger.Debug("enter handler")
	defer logger.Debug("exit handler")

	body := req.Body
	defer func() {
		if err := body.Close(); err != nil {
			logger.Error("failed to close body", zap.Error(err))
		}
	}()

	var bodyBytes []byte
	if _, err := body.Read(bodyBytes); err != nil {
		logger.Error("failed to read request body", zap.Error(err))
		render(w, http.StatusInternalServerError)
		return
	}

	var createReq *model.CreateRequest
	if err := json.Unmarshal(bodyBytes, createReq); err != nil {
		logger.Error("failed to unmarshal request body", zap.Error(err))
		render(w, http.StatusBadRequest)
		return
	}

	if err := app.service.Create(ctx, createReq); err != nil {
		logger.Error("service failed", zap.Error(err))
		renderServiceError(w, err)
		return
	}
}

func (app *App) updateAnswer(w http.ResponseWriter, req *http.Request) {

}

func (app *App) deleteAnswer(w http.ResponseWriter, req *http.Request) {

}

func (app *App) getAnswer(w http.ResponseWriter, req *http.Request) {

}

func (app *App) getAnswerEvents(w http.ResponseWriter, req *http.Request) {

}
