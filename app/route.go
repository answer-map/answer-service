package app

import "net/http"

const (
	basePath   = "/api/v1"
	answerPath = "/answers"
)

func (app *App) registerRoutes() {
	apiRouter := app.router.Path(basePath + answerPath).Subrouter()

	apiRouter.Methods(http.MethodPost).HandlerFunc(app.createAnswer)

	apiRouter.Methods(http.MethodPut).HandlerFunc(app.updateAnswer)
}
