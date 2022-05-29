package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/answer-map/answer-repository/entity"
	"github.com/answer-map/answer-service/model"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//go:generate mockgen -source .\service\answer.go -destination .\mock_service\answer.go

var (
	nullString = null.NewString("", false)
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

	if !exists {
		tx.Rollback()

		return ErrConflict
	}

	answerMap := entity.AnswerMap{
		AnswerKey:   req.AnswerKey,
		AnswerValue: null.StringFrom(req.AnswerValue),
	}

	if _, err = answerMap.Update(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	answerEvent := entity.AnswerEvent{
		EventID:        uuid.NewString(),
		EventType:      entity.EventTypeUpdate,
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

func (s *answerService) Delete(ctx context.Context, req *model.DeleteRequest) error {
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

	if !exists {
		tx.Rollback()

		return ErrConflict
	}

	answerMap := entity.AnswerMap{
		AnswerKey:   req.AnswerKey,
		AnswerValue: nullString,
	}

	if _, err = answerMap.Update(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	answerEvent := entity.AnswerEvent{
		EventID:        uuid.NewString(),
		EventType:      entity.EventTypeDelete,
		EventTimestamp: time.Now(),
		AnswerKey:      req.AnswerKey,
		AnswerValue:    nullString,
	}

	if err := answerEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}

func (s *answerService) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, ErrBadRequest
	}

	answerMap, err := entity.AnswerMaps(entity.AnswerMapWhere.AnswerKey.EQ(req.AnswerKey), entity.AnswerMapWhere.AnswerValue.IsNotNull()).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrResourceNotFound
		}

		return nil, err
	}

	return &model.AnswerMap{AnswerKey: answerMap.AnswerKey, AnswerValue: answerMap.AnswerValue.String}, nil
}

func (s *answerService) GetHistory(ctx context.Context, req *model.GetHistoryRequest) (*model.GetHistoryResponse, error) {
	answerEvents, err := entity.AnswerEvents(entity.AnswerEventWhere.EventTimestamp.GTE(req.MinimumEventTimestamp), qm.OrderBy(entity.AnswerEventTableColumns.EventTimestamp), qm.Limit(int(req.PageSize))).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if len(answerEvents) == 0 {
		return nil, ErrResourceNotFound
	}

	var res model.GetHistoryResponse

	for i := range answerEvents {
		res = append(res, model.FromEntity(answerEvents[i]))
	}

	return &res, nil
}
