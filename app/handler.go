package app

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/answer-map/answer-service/model"
	"go.uber.org/zap"
)

func (app *App) readAnswerMap(w http.ResponseWriter, req *http.Request) (context.Context, *model.AnswerMap) {
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

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		logger.Error("failed to read request body", zap.Error(err))
		render(w, http.StatusInternalServerError)
		return ctx, nil
	}

	var answerMap model.AnswerMap
	if err := json.Unmarshal(bodyBytes, &answerMap); err != nil {
		logger.Error("failed to unmarshal request body", zap.Error(err))
		render(w, http.StatusBadRequest)
		return ctx, nil
	}

	return ctx, &answerMap
}

func (app *App) readAnswerMapKey(w http.ResponseWriter, req *http.Request) (context.Context, *model.AnswerMapKey) {
	ctx := req.Context()
	logger := app.logger
	logger.Debug("enter handler")
	defer logger.Debug("exit handler")

	answerKey := req.URL.Query().Get("answerKey")
	if len(answerKey) == 0 {
		logger.Error("failed to read answer key")
		render(w, http.StatusBadRequest)
		return ctx, nil
	}

	return ctx, &model.AnswerMapKey{AnswerKey: answerKey}
}

func (app *App) createAnswer(w http.ResponseWriter, req *http.Request) {
	ctx, answerMap := app.readAnswerMap(w, req)
	if answerMap == nil {
		return
	}

	if err := app.service.Create(ctx, answerMap); err != nil {
		app.logger.Error("service failed", zap.Error(err))
		renderServiceError(w, err)
		return
	}
}

func (app *App) updateAnswer(w http.ResponseWriter, req *http.Request) {
	ctx, answerMap := app.readAnswerMap(w, req)
	if answerMap == nil {
		return
	}

	if err := app.service.Update(ctx, answerMap); err != nil {
		app.logger.Error("service failed", zap.Error(err))
		renderServiceError(w, err)
		return
	}
}

func (app *App) deleteAnswer(w http.ResponseWriter, req *http.Request) {
	ctx, answerKey := app.readAnswerMapKey(w, req)
	if answerKey == nil {
		return
	}

	if err := app.service.Delete(ctx, answerKey); err != nil {
		app.logger.Error("service failed", zap.Error(err))
		renderServiceError(w, err)
		return
	}
}

func (app *App) getAnswer(w http.ResponseWriter, req *http.Request) {
	ctx, answerKey := app.readAnswerMapKey(w, req)
	if answerKey == nil {
		return
	}

	answerMap, err := app.service.Get(ctx, answerKey)
	if err != nil {
		app.logger.Error("service failed", zap.Error(err))
		renderServiceError(w, err)
		return
	}

	renderMarshalJSON(w, http.StatusOK, answerMap)
}

func (app *App) getAnswerEvents(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logger := app.logger
	logger.Debug("enter handler")
	defer logger.Debug("exit handler")

	query, err := model.NewGetHistoryRequest(req.URL.Query())
	if err != nil {
		logger.Error("failed to get query params", zap.Error(err))
		render(w, http.StatusBadRequest)
		return
	}

	answerEvents, err := app.service.GetHistory(ctx, query)
	if err != nil {
		app.logger.Error("service failed", zap.Error(err))
		renderServiceError(w, err)
		return
	}

	renderMarshalJSON(w, http.StatusOK, answerEvents)
}
