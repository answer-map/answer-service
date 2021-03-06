package app

import "net/http"

const (
	basePath        = "/api/v1"
	answerUserPath  = basePath + "/users"
	answerPath      = basePath + "/answers"
	answerEventPath = basePath + "/events"
)

func (app *App) registerRoutes() {
	app.router.HandleFunc(answerUserPath, app.createAnswerUser).Methods(http.MethodPost)

	app.router.HandleFunc(answerPath, app.createAnswer).Methods(http.MethodPost)
	app.router.HandleFunc(answerPath, app.updateAnswer).Methods(http.MethodPut)
	app.router.HandleFunc(answerPath, app.deleteAnswer).Methods(http.MethodDelete)
	app.router.HandleFunc(answerPath, app.getAnswer).Methods(http.MethodGet)

	app.router.HandleFunc(answerEventPath, app.getAnswerEvents).Methods(http.MethodGet)
}
