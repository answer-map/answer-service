package app

import "net/http"

const (
	basePath        = "/api/v1"
	answerPath      = basePath + "/answers"
	answerKeyPath   = answerPath + "/{answerKey}"
	answerEventPath = basePath + "/answerEvents"
)

func (app *App) registerRoutes() {
	app.router.HandleFunc(answerPath, app.createAnswer).Methods(http.MethodPost)
	app.router.HandleFunc(answerPath, app.updateAnswer).Methods(http.MethodPut)

	app.router.HandleFunc(answerKeyPath, app.deleteAnswer).Methods(http.MethodDelete)
	app.router.HandleFunc(answerKeyPath, app.getAnswer).Methods(http.MethodGet)

	app.router.HandleFunc(answerEventPath, app.getAnswerEvents).Methods(http.MethodGet)
}
