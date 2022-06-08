package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/answer-map/answer-service/entity"
	"github.com/answer-map/answer-service/model"
	"github.com/friendsofgo/errors"
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
	CreateUser(ctx context.Context, req *model.CreateUserRequest) error
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

func (s *answerService) CreateUser(ctx context.Context, req *model.CreateUserRequest) error {
	if err := req.Validate(); err != nil {
		return ErrBadRequest
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	exists, err := entity.AnswerUsers(entity.AnswerUserWhere.UserID.EQ(req.UserID)).Exists(ctx, tx)
	if err != nil {
		return err
	}

	if exists {
		tx.Rollback()

		return ErrConflict
	}

	answerUser := entity.AnswerUser{UserID: req.UserID}

	if err := answerUser.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (s *answerService) Create(ctx context.Context, req *model.CreateRequest) error {
	if err := req.Validate(); err != nil {
		return ErrBadRequest
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	exists, err := entity.AnswerMaps(entity.AnswerMapWhere.UserID.EQ(req.UserID), entity.AnswerMapWhere.AnswerKey.EQ(req.AnswerKey), entity.AnswerMapWhere.AnswerValue.IsNotNull()).Exists(ctx, tx)
	if err != nil {
		return err
	}

	if exists {
		tx.Rollback()

		return ErrConflict
	}

	answerMap, err := entity.AnswerMaps(entity.AnswerMapWhere.UserID.EQ(req.UserID), entity.AnswerMapWhere.AnswerKey.EQ(req.AnswerKey), entity.AnswerMapWhere.AnswerValue.IsNull()).One(ctx, tx)
	if err == nil {
		if _, err := answerMap.Update(ctx, tx, boil.Whitelist(entity.AnswerMapColumns.AnswerValue)); err != nil {
			tx.Rollback()

			return err
		}
	} else {
		if !errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()

			return err
		}

		answerMap = &entity.AnswerMap{
			MapID:       uuid.NewString(),
			UserID:      req.UserID,
			AnswerKey:   req.AnswerKey,
			AnswerValue: null.StringFrom(req.AnswerValue),
		}

		if err := answerMap.Insert(ctx, tx, boil.Infer()); err != nil {
			tx.Rollback()

			return err
		}
	}

	answerEvent := entity.AnswerEvent{
		EventID:        uuid.NewString(),
		EventType:      entity.EventTypeCreate,
		EventTimestamp: time.Now(),
		MapID:          answerMap.MapID,
		AnswerValue:    null.StringFrom(req.AnswerValue),
	}

	if err := answerEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (s *answerService) Update(ctx context.Context, req *model.UpdateRequest) error {
	if err := req.Validate(); err != nil {
		return ErrBadRequest
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	answerMap, err := entity.AnswerMaps(entity.AnswerMapWhere.UserID.EQ(req.UserID), entity.AnswerMapWhere.AnswerKey.EQ(req.AnswerKey), entity.AnswerMapWhere.AnswerValue.IsNotNull()).One(ctx, tx)
	if err != nil {
		tx.Rollback()

		if errors.Is(err, sql.ErrNoRows) {
			return ErrConflict
		}

		return err
	}

	answerMap.AnswerValue = null.StringFrom(req.AnswerValue)

	if _, err = answerMap.Update(ctx, tx, boil.Whitelist(entity.AnswerMapColumns.AnswerValue)); err != nil {
		tx.Rollback()

		return err
	}

	answerEvent := entity.AnswerEvent{
		EventID:        uuid.NewString(),
		EventType:      entity.EventTypeUpdate,
		EventTimestamp: time.Now(),
		MapID:          answerMap.MapID,
		AnswerValue:    null.StringFrom(req.AnswerValue),
	}

	if err := answerEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (s *answerService) Delete(ctx context.Context, req *model.DeleteRequest) error {
	if err := req.Validate(); err != nil {
		return ErrBadRequest
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	answerMap, err := entity.AnswerMaps(entity.AnswerMapWhere.UserID.EQ(req.UserID), entity.AnswerMapWhere.AnswerKey.EQ(req.AnswerKey), entity.AnswerMapWhere.AnswerValue.IsNotNull()).One(ctx, tx)
	if err != nil {
		tx.Rollback()

		if errors.Is(err, sql.ErrNoRows) {
			return ErrConflict
		}

		return err
	}

	answerMap.AnswerValue = nullString

	if _, err = answerMap.Update(ctx, tx, boil.Whitelist(entity.AnswerMapColumns.AnswerValue)); err != nil {
		tx.Rollback()

		return err
	}

	answerEvent := entity.AnswerEvent{
		EventID:        uuid.NewString(),
		EventType:      entity.EventTypeDelete,
		EventTimestamp: time.Now(),
		MapID:          answerMap.MapID,
		AnswerValue:    nullString,
	}

	if err := answerEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (s *answerService) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, ErrBadRequest
	}

	answerMap, err := entity.AnswerMaps(entity.AnswerMapWhere.UserID.EQ(req.UserID), entity.AnswerMapWhere.AnswerKey.EQ(req.AnswerKey), entity.AnswerMapWhere.AnswerValue.IsNotNull()).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrResourceNotFound
		}

		return nil, err
	}

	return &model.AnswerMap{UserID: answerMap.UserID, AnswerKey: answerMap.AnswerKey, AnswerValue: answerMap.AnswerValue.String}, nil
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
