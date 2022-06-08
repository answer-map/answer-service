package service_test

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/answer-map/answer-service/entity"
	"github.com/answer-map/answer-service/model"
	"github.com/answer-map/answer-service/service"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

const (
	sqlCountAnswerUsers              string = `SELECT COUNT(*) FROM "answer_user" WHERE ("answer_user"."user_id" = $1) LIMIT 1;`
	sqlInsertAnswerUser              string = `INSERT INTO "answer_user" ("user_id") VALUES ($1)`
	sqlCountAnswerMaps               string = `SELECT COUNT(*) FROM "answer_map" WHERE ("answer_map"."user_id" = $1) AND ("answer_map"."answer_key" = $2) AND ("answer_map"."answer_value" is not null) LIMIT 1;`
	sqlSelectAnswerMapWithValNotNull string = `SELECT "answer_map".* FROM "answer_map" WHERE ("answer_map"."user_id" = $1) AND ("answer_map"."answer_key" = $2) AND ("answer_map"."answer_value" is not null) LIMIT 1;`
	sqlSelectAnswerMapWithValNull    string = `SELECT "answer_map".* FROM "answer_map" WHERE ("answer_map"."user_id" = $1) AND ("answer_map"."answer_key" = $2) AND ("answer_map"."answer_value" is null) LIMIT 1;`
	sqlInsertAnswerMap               string = `INSERT INTO "answer_map" ("map_id","user_id","answer_key","answer_value") VALUES ($1,$2,$3,$4)`
	sqlUpdateAnswerMap               string = `UPDATE "answer_map" SET "answer_value"=$1 WHERE "map_id"=$2`
	sqlInsertAnswerEvent             string = `INSERT INTO "answer_event" ("event_id","event_type","event_timestamp","map_id","answer_value") VALUES ($1,$2,$3,$4,$5)`
	sqlInsertAnswerEvent2            string = `INSERT INTO "answer_event" ("event_id","event_type","event_timestamp","map_id") VALUES ($1,$2,$3,$4) RETURNING "answer_value"`
	sqlSelectAnswerEvent             string = `SELECT "answer_event".* FROM "answer_event" WHERE ("answer_event"."event_timestamp" >= $1) ORDER BY answer_event.event_timestamp LIMIT 2;`
	argCount                         string = "count"
	argAnswerValue                   string = "answer_value"
)

var (
	uuid1         = uuid.NewString()
	uuid2         = uuid.NewString()
	uuid3         = uuid.NewString()
	answerUser1   = &model.AnswerUser{UserID: uuid2}
	answerMap1    = &model.AnswerMap{UserID: uuid2, AnswerKey: "name", AnswerValue: ""}
	answerMap2    = &model.AnswerMap{UserID: uuid2, AnswerKey: "", AnswerValue: "harry"}
	answerMap3    = &model.AnswerMap{UserID: uuid2, AnswerKey: "name", AnswerValue: "harry"}
	answerMapKey1 = &model.AnswerMapKey{UserID: uuid2, AnswerKey: ""}
	answerMapKey2 = &model.AnswerMapKey{UserID: uuid2, AnswerKey: "name"}
	nullString    = null.NewString("", false)
	getHistoryReq = &model.GetHistoryRequest{MinimumEventTimestamp: time.Date(2022, time.May, 29, 0, 0, 0, 0, time.UTC), PageSize: 2}
)

func getAnswerMapRows(mapID string, answerMap *model.AnswerMap) *sqlmock.Rows {
	return sqlmock.NewRows([]string{entity.AnswerMapColumns.MapID, entity.AnswerMapColumns.UserID, entity.AnswerMapColumns.AnswerKey, entity.AnswerMapColumns.AnswerValue}).AddRow(mapID, answerMap.UserID, answerMap.AnswerKey, answerMap.AnswerValue)
}

func Test_answerService_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *model.CreateUserRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name: "user id is used",
			args: args{
				ctx: context.Background(),
				req: answerUser1,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlCountAnswerUsers)).WithArgs(answerUser1.UserID).WillReturnRows(sqlmock.NewRows([]string{argCount}).AddRow(1))
				mock.ExpectRollback()
			},
			wantErr: service.ErrConflict,
		},
		{
			name: "user id is new",
			args: args{
				ctx: context.Background(),
				req: answerUser1,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlCountAnswerUsers)).WithArgs(answerUser1.UserID).WillReturnRows(sqlmock.NewRows([]string{argCount}).AddRow(0))
				mock.ExpectExec(regexp.QuoteMeta(sqlInsertAnswerUser)).WithArgs(answerUser1.UserID).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("answerServiceServer db init error = %v", err)
			}
			defer db.Close()

			if tt.expectations == nil {
				t.Fatal("answerServiceServer expectations not set")
			}

			tt.expectations(mock)

			server := service.NewAnswerService(db)

			if err := server.CreateUser(tt.args.ctx, tt.args.req); err != tt.wantErr {
				t.Errorf("answerServiceServer.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_answerService_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		req *model.CreateRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name: "key is name, value is empty",
			args: args{
				ctx: context.Background(),
				req: answerMap1,
			},
			expectations: func(mock sqlmock.Sqlmock) {},
			wantErr:      service.ErrBadRequest,
		},
		{
			name: "key is empty, value is harry",
			args: args{
				ctx: context.Background(),
				req: answerMap2,
			},
			expectations: func(mock sqlmock.Sqlmock) {},
			wantErr:      service.ErrBadRequest,
		},
		{
			name: "key is name, value is harry, there is already a name",
			args: args{
				ctx: context.Background(),
				req: answerMap3,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlCountAnswerMaps)).WithArgs(answerMap3.UserID, answerMap3.AnswerKey).WillReturnRows(sqlmock.NewRows([]string{argCount}).AddRow(1))
				mock.ExpectRollback()
			},
			wantErr: service.ErrConflict,
		},
		{
			name: "key is name, value is harry, update",
			args: args{
				ctx: context.Background(),
				req: answerMap3,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlCountAnswerMaps)).WithArgs(answerMap3.UserID, answerMap3.AnswerKey).WillReturnRows(sqlmock.NewRows([]string{argCount}).AddRow(0))
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNull)).WithArgs(answerMap3.UserID, answerMap3.AnswerKey).WillReturnRows(getAnswerMapRows(uuid3, answerMap3))
				mock.ExpectExec(regexp.QuoteMeta(sqlUpdateAnswerMap)).WithArgs(answerMap3.AnswerValue, uuid3).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(regexp.QuoteMeta(sqlInsertAnswerEvent)).WithArgs(sqlmock.AnyArg(), entity.EventTypeCreate, sqlmock.AnyArg(), uuid3, answerMap3.AnswerValue).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "key is name, value is harry, insert",
			args: args{
				ctx: context.Background(),
				req: answerMap3,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlCountAnswerMaps)).WithArgs(answerMap3.UserID, answerMap3.AnswerKey).WillReturnRows(sqlmock.NewRows([]string{argCount}).AddRow(0))
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNull)).WithArgs(answerMap3.UserID, answerMap3.AnswerKey).WillReturnError(sql.ErrNoRows)
				mock.ExpectExec(regexp.QuoteMeta(sqlInsertAnswerMap)).WithArgs(sqlmock.AnyArg(), answerMap3.UserID, answerMap3.AnswerKey, answerMap3.AnswerValue).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(regexp.QuoteMeta(sqlInsertAnswerEvent)).WithArgs(sqlmock.AnyArg(), entity.EventTypeCreate, sqlmock.AnyArg(), sqlmock.AnyArg(), answerMap3.AnswerValue).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("answerServiceServer db init error = %v", err)
			}
			defer db.Close()

			if tt.expectations == nil {
				t.Fatal("answerServiceServer expectations not set")
			}

			tt.expectations(mock)

			server := service.NewAnswerService(db)

			if err := server.Create(tt.args.ctx, tt.args.req); err != tt.wantErr {
				t.Errorf("answerServiceServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_answerService_Update(t *testing.T) {
	type args struct {
		ctx context.Context
		req *model.UpdateRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name: "key is name, value is empty",
			args: args{
				ctx: context.Background(),
				req: answerMap1,
			},
			expectations: func(mock sqlmock.Sqlmock) {},
			wantErr:      service.ErrBadRequest,
		},
		{
			name: "key is empty, value is harry",
			args: args{
				ctx: context.Background(),
				req: answerMap2,
			},
			expectations: func(mock sqlmock.Sqlmock) {},
			wantErr:      service.ErrBadRequest,
		},
		{
			name: "key is name, value is harry, there is not a name",
			args: args{
				ctx: context.Background(),
				req: answerMap3,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNotNull)).WithArgs(answerMap3.UserID, answerMap3.AnswerKey).WillReturnError(sql.ErrNoRows)
				mock.ExpectRollback()
			},
			wantErr: service.ErrConflict,
		},
		{
			name: "key is name, value is harry, complete",
			args: args{
				ctx: context.Background(),
				req: answerMap3,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNotNull)).WithArgs(answerMap3.UserID, answerMap3.AnswerKey).WillReturnRows(getAnswerMapRows(uuid3, answerMap3))
				mock.ExpectExec(regexp.QuoteMeta(sqlUpdateAnswerMap)).WithArgs(answerMap3.AnswerValue, uuid3).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(regexp.QuoteMeta(sqlInsertAnswerEvent)).WithArgs(sqlmock.AnyArg(), entity.EventTypeUpdate, sqlmock.AnyArg(), uuid3, answerMap3.AnswerValue).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("answerServiceServer db init error = %v", err)
			}
			defer db.Close()

			if tt.expectations == nil {
				t.Fatal("answerServiceServer expectations not set")
			}

			tt.expectations(mock)

			server := service.NewAnswerService(db)

			if err := server.Update(tt.args.ctx, tt.args.req); err != tt.wantErr {
				t.Errorf("answerService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_answerService_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		req *model.DeleteRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name: "key is empty",
			args: args{
				ctx: context.Background(),
				req: answerMapKey1,
			},
			expectations: func(mock sqlmock.Sqlmock) {},
			wantErr:      service.ErrBadRequest,
		},
		{
			name: "key is name, there is not a name",
			args: args{
				ctx: context.Background(),
				req: answerMapKey2,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNotNull)).WithArgs(answerMapKey2.UserID, answerMapKey2.AnswerKey).WillReturnError(sql.ErrNoRows)
				mock.ExpectRollback()
			},
			wantErr: service.ErrConflict,
		},
		{
			name: "key is name, there is a name, complete",
			args: args{
				ctx: context.Background(),
				req: answerMapKey2,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNotNull)).WithArgs(answerMapKey2.UserID, answerMapKey2.AnswerKey).WillReturnRows(getAnswerMapRows(uuid3, answerMap3))
				mock.ExpectExec(regexp.QuoteMeta(sqlUpdateAnswerMap)).WithArgs(nullString, uuid3).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(regexp.QuoteMeta(sqlInsertAnswerEvent2)).WithArgs(sqlmock.AnyArg(), entity.EventTypeDelete, sqlmock.AnyArg(), uuid3).WillReturnRows(sqlmock.NewRows([]string{argAnswerValue}).AddRow(nullString))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("answerServiceServer db init error = %v", err)
			}
			defer db.Close()

			if tt.expectations == nil {
				t.Fatal("answerServiceServer expectations not set")
			}

			tt.expectations(mock)

			server := service.NewAnswerService(db)

			if err := server.Delete(tt.args.ctx, tt.args.req); err != tt.wantErr {
				t.Errorf("answerService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_answerService_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		req *model.GetRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		want         *model.GetResponse
		wantErr      error
	}{
		{
			name: "key is empty",
			args: args{
				ctx: context.Background(),
				req: answerMapKey1,
			},
			expectations: func(mock sqlmock.Sqlmock) {},
			wantErr:      service.ErrBadRequest,
		},
		{
			name: "key is name, there is not a name",
			args: args{
				ctx: context.Background(),
				req: answerMapKey2,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNotNull)).WithArgs(answerMapKey2.UserID, answerMapKey2.AnswerKey).WillReturnError(sql.ErrNoRows)
			},
			wantErr: service.ErrResourceNotFound,
		},
		{
			name: "key is name, there is a name, complete",
			args: args{
				ctx: context.Background(),
				req: answerMapKey2,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerMapWithValNotNull)).WithArgs(answerMapKey2.UserID, answerMapKey2.AnswerKey).WillReturnRows(getAnswerMapRows(uuid3, answerMap3))
			},
			want: answerMap3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("answerServiceServer db init error = %v", err)
			}
			defer db.Close()

			if tt.expectations == nil {
				t.Fatal("answerServiceServer expectations not set")
			}

			tt.expectations(mock)

			server := service.NewAnswerService(db)

			got, err := server.Get(tt.args.ctx, tt.args.req)
			if err != tt.wantErr {
				t.Errorf("answerService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_answerService_GetHistory(t *testing.T) {
	type args struct {
		ctx context.Context
		req *model.GetHistoryRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		want         *model.GetHistoryResponse
		wantErr      error
	}{
		{
			name: "no rows",
			args: args{
				ctx: context.Background(),
				req: getHistoryReq,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerEvent)).WithArgs(getHistoryReq.MinimumEventTimestamp).WillReturnRows(sqlmock.NewRows([]string{
					entity.AnswerEventColumns.EventID,
					entity.AnswerEventColumns.EventTimestamp,
					entity.AnswerEventColumns.EventType,
					entity.AnswerEventColumns.MapID,
					entity.AnswerEventColumns.AnswerValue}))
			},
			wantErr: service.ErrResourceNotFound,
		},
		{
			name: "one row",
			args: args{
				ctx: context.Background(),
				req: getHistoryReq,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAnswerEvent)).WithArgs(getHistoryReq.MinimumEventTimestamp).WillReturnRows(sqlmock.NewRows([]string{
					entity.AnswerEventColumns.EventID,
					entity.AnswerEventColumns.EventTimestamp,
					entity.AnswerEventColumns.EventType,
					entity.AnswerEventColumns.MapID,
					entity.AnswerEventColumns.AnswerValue}).AddRow(
					uuid1,
					getHistoryReq.MinimumEventTimestamp,
					entity.EventTypeCreate,
					uuid3,
					null.StringFrom(answerMap3.AnswerValue),
				))
			},
			want: &[]model.AnswerEvent{
				{
					EventID:        uuid1,
					EventTimestamp: getHistoryReq.MinimumEventTimestamp,
					EventType:      entity.EventTypeCreate,
					MapID:          uuid3,
					AnswerValue:    null.StringFrom(answerMap3.AnswerValue),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("answerServiceServer db init error = %v", err)
			}
			defer db.Close()

			if tt.expectations == nil {
				t.Fatal("answerServiceServer expectations not set")
			}

			tt.expectations(mock)

			server := service.NewAnswerService(db)

			got, err := server.GetHistory(tt.args.ctx, tt.args.req)
			if err != tt.wantErr {
				t.Errorf("answerService.GetHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerService.GetHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}
