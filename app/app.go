package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/answer-map/answer-service/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type App struct {
	address string
	logger  *zap.Logger
	db      *sql.DB
	service service.AnswerService
	router  *mux.Router
}

func NewApp(config *Config) (*App, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config, error = %v", err)
	}

	var app App
	app.address = config.HTTP.Address()

	logger, err := config.ZapLogger.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to start zap logger, error = %v", err)
	}
	app.logger = logger

	db, err := sql.Open("postgres", config.AnswerDataBase.DataSource())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection, error = %v", err)
	}
	app.db = db

	app.service = service.NewAnswerService(db)

	app.router = mux.NewRouter()
	app.registerRoutes()

	return &app, nil
}

func (app *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := app.logger

	httpServer := &http.Server{
		Addr:    app.address,
		Handler: app.router,
	}

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sigterm := <-termChan
		logger.Info("shutdown process initiated", zap.Any("sigterm", sigterm))
		httpServer.Shutdown(ctx)
	}()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logger.Error("http server closed", zap.Error(err))
			}
			logger.Info("http server shut down")
		}

		if err := app.db.Close(); err != nil {
			logger.Error("db connection closed", zap.Error(err))
		}
		logger.Info("db connection shut down")

		cancel()
	}()

	logger.Info("app started")
	<-ctx.Done()
	logger.Info("app shut down")
}
