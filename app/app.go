package app

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/lib/pq"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	answer_protobuf "github.com/answer-map/answer-protobuf"
	"github.com/answer-map/answer-service/service_server"
)

type App struct {
	grpcAddress string
	logger      *zap.Logger
	grpcServer  *grpc.Server
}

func NewApp(config *Config) (*App, error) {
	var app App
	app.grpcAddress = config.GRPCHTTP.Address()

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config, error = %v", err)
	}

	logger, err := config.ZapLogger.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to start zap logger, error = %v", err)
	}

	app.logger = logger

	// Open handle to database like normal
	db, err := sql.Open("postgres", config.AnswerDataBase.DataSource())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection, error = %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	answer_protobuf.RegisterAnswerServiceServer(grpcServer, service_server.NewAnswerServiceServer(db))

	app.grpcServer = grpcServer

	return &app, nil
}

func (app *App) Run() error {
	lis, err := net.Listen("tcp", app.grpcAddress)
	if err != nil {
		return fmt.Errorf("failed to create listener, error = %v", err)
	}

	mux := cmux.New(lis)
	grpcL := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	go func() {
		sErr := mux.Serve()
		if sErr != nil {
			app.logger.Fatal("failed to serve cmux", zap.Error(err))
		}
	}()

	app.logger.With(zap.Any("address", grpcL.Addr())).Info("serving grpc")

	return app.grpcServer.Serve(grpcL)
}
