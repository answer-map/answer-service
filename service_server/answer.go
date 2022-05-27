package service_server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	answer_protobuf "github.com/answer-map/answer-protobuf"
	"github.com/answer-map/answer-repository/entities"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type answerServiceServer struct {
	db *sql.DB
	answer_protobuf.UnimplementedAnswerServiceServer
}

func NewAnswerServiceServer(db *sql.DB) *answerServiceServer {
	return &answerServiceServer{
		db: db,
	}
}

func (s *answerServiceServer) Create(ctx context.Context, req *answer_protobuf.CreateRequest) (*answer_protobuf.CreateResponse, error) {
	statusCodeRes := func(code uint32) (*answer_protobuf.CreateResponse, error) {
		return &answer_protobuf.CreateResponse{StatusCode: code}, nil
	}

	if req.Answer.Value == "" {
		return statusCodeRes(400)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	exists, err := entities.AnswerMaps(entities.AnswerMapWhere.AnswerKey.EQ(req.Answer.Key), entities.AnswerMapWhere.AnswerValue.IsNotNull()).Exists(ctx, tx)
	if err != nil {
		return nil, err
	}

	if exists {
		tx.Rollback()

		return statusCodeRes(409)
	}

	answerMap := entities.AnswerMap{
		AnswerKey:   req.Answer.Key,
		AnswerValue: null.StringFrom(req.Answer.Value),
	}

	if err := answerMap.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return nil, err
	}

	answerEvent := entities.AnswerEvent{
		EventID:        uuid.NewString(),
		EventType:      entities.EventTypeCreate,
		EventTimestamp: time.Now(),
		AnswerKey:      req.Answer.Key,
		AnswerValue:    null.StringFrom(req.Answer.Value),
	}

	if err := answerEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return nil, err
	}

	return statusCodeRes(200)
}

func (s *answerServiceServer) Update(context.Context, *answer_protobuf.UpdateRequest) (*answer_protobuf.UpdateResponse, error) {
	return nil, fmt.Errorf("TODO: unimplemented")
}

func (s *answerServiceServer) Delete(context.Context, *answer_protobuf.DeleteRequest) (*answer_protobuf.DeleteResponse, error) {
	return nil, fmt.Errorf("TODO: unimplemented")
}

func (s *answerServiceServer) Get(context.Context, *answer_protobuf.GetRequest) (*answer_protobuf.GetResponse, error) {
	return nil, fmt.Errorf("TODO: unimplemented")
}

func (s *answerServiceServer) GetHistory(context.Context, *answer_protobuf.GetRequest) (*answer_protobuf.GetResponse, error) {
	return nil, fmt.Errorf("TODO: unimplemented")
}
