package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/answer-map/answer-repository/entity"
	"github.com/answer-map/answer-service/model"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AnswerService interface {
	Create(ctx context.Context, req *model.CreateRequest) error
	Update(ctx context.Context, req *model.UpdateRequest) error
	Delete(ctx context.Context, req *model.DeleteRequest) error
	Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error)
	GetHistory(ctx context.Context, req *model.GetHistoryRequest) (*model.GetHistoryResponse, error)
}

type answerService struct {
	db *sql.DB
}

func NewAnswerService(db *sql.DB) *answerService {
	return &answerService{
		db: db,
	}
}

func (s *answerService) Create(ctx context.Context, req *model.CreateRequest) error {
	if err := req.Validate(); err != nil {
		return ErrBadRequest
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	exists, err := entity.AnswerMaps(entity.AnswerMapWhere.AnswerKey.EQ(req.AnswerKey), entity.AnswerMapWhere.AnswerValue.IsNotNull()).Exists(ctx, tx)
	if err != nil {
		return err
	}

	if exists {
		tx.Rollback()

		return ErrConflict
	}

	answerMap := entity.AnswerMap{
		AnswerKey:   req.AnswerKey,
		AnswerValue: null.StringFrom(req.AnswerValue),
	}

	if err := answerMap.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	answerEvent := entity.AnswerEvent{
		EventID:        uuid.NewString(),
		EventType:      entity.EventTypeCreate,
		EventTimestamp: time.Now(),
		AnswerKey:      req.AnswerKey,
		AnswerValue:    null.StringFrom(req.AnswerValue),
	}

	if err := answerEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}

func (s *answerService) Update(ctx context.Context, req *model.UpdateRequest) error {
	return fmt.Errorf("TODO: unimplemented")
}

func (s *answerService) Delete(ctx context.Context, req *model.DeleteRequest) error {
	return fmt.Errorf("TODO: unimplemented")
}

func (s *answerService) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	return nil, fmt.Errorf("TODO: unimplemented")
}

func (s *answerService) GetHistory(ctx context.Context, req *model.GetHistoryRequest) (*model.GetHistoryResponse, error) {
	return nil, fmt.Errorf("TODO: unimplemented")
}
