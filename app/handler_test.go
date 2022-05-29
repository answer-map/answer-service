package app

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"context"

	"github.com/answer-map/answer-repository/entity"
	"github.com/answer-map/answer-service/mock_service"
	"github.com/answer-map/answer-service/model"
	"github.com/answer-map/answer-service/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

const (
	answerMapString1        string = `{"answerKey":"name","answerValue":"lewis"}`
	answerKeyQueryString1   string = `?answerKey=name`
	answerEventQueryString1 string = `?minimumEventTimestamp=2022-05-28T00:00:00Z&pageSize=2`
)

var (
	answerMap1             = &model.AnswerMap{AnswerKey: "name", AnswerValue: "lewis"}
	answerMapKey1          = &model.AnswerMapKey{AnswerKey: "name"}
	uuid1                  = uuid.NewString()
	time1                  = time.Date(2022, 5, 28, 0, 0, 0, 0, time.UTC)
	getHistoryRequest1     = &model.GetHistoryRequest{MinimumEventTimestamp: time1, PageSize: 2}
	answerEvent1           = &model.AnswerEvent{EventID: uuid1, EventType: entity.EventTypeCreate, EventTimestamp: time1, AnswerKey: "name", AnswerValue: null.StringFrom("lewis")}
	answerEventList1       = &[]model.AnswerEvent{*answerEvent1}
	answerEventListString1 = fmt.Sprintf(`[{"eventID":"%s","eventType":"create","eventTimestamp":"%s","answerKey":"name","answerValue":"lewis"}]`, uuid1, time1.Format(time.RFC3339))
)

func TestApp_handlers(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name         string
		args         args
		expectations func(ctx context.Context, mock *mock_service.MockAnswerService)
		handler      func(app *App, w http.ResponseWriter, req *http.Request)
		wantStatus   int
		wantBody     string
	}{
		{
			name: "create answer map, bad request",
			args: args{
				method: http.MethodPost,
				url:    answerPath,
				body:   strings.NewReader(answerMapString1),
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Create(ctx, answerMap1).Return(service.ErrBadRequest)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.createAnswer(w, req)
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "create answer map, conflict",
			args: args{
				method: http.MethodPost,
				url:    answerPath,
				body:   strings.NewReader(answerMapString1),
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Create(ctx, answerMap1).Return(service.ErrConflict)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.createAnswer(w, req)
			},
			wantStatus: http.StatusConflict,
		},
		{
			name: "create answer map, ok",
			args: args{
				method: http.MethodPost,
				url:    answerPath,
				body:   strings.NewReader(answerMapString1),
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Create(ctx, answerMap1).Return(nil)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.createAnswer(w, req)
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "update answer map, bad request",
			args: args{
				method: http.MethodPut,
				url:    answerPath,
				body:   strings.NewReader(answerMapString1),
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Update(ctx, answerMap1).Return(service.ErrBadRequest)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.updateAnswer(w, req)
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "update answer map, conflict",
			args: args{
				method: http.MethodPut,
				url:    answerPath,
				body:   strings.NewReader(answerMapString1),
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Update(ctx, answerMap1).Return(service.ErrConflict)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.updateAnswer(w, req)
			},
			wantStatus: http.StatusConflict,
		},
		{
			name: "update answer map, ok",
			args: args{
				method: http.MethodPut,
				url:    answerPath,
				body:   strings.NewReader(answerMapString1),
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Update(ctx, answerMap1).Return(nil)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.updateAnswer(w, req)
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "delete answer map, bad request",
			args: args{
				method: http.MethodDelete,
				url:    answerPath,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.deleteAnswer(w, req)
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "delete answer map, conflict",
			args: args{
				method: http.MethodDelete,
				url:    answerPath + answerKeyQueryString1,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Delete(ctx, answerMapKey1).Return(service.ErrConflict)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.deleteAnswer(w, req)
			},
			wantStatus: http.StatusConflict,
		},
		{
			name: "delete answer map, ok",
			args: args{
				method: http.MethodDelete,
				url:    answerPath + answerKeyQueryString1,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Delete(ctx, answerMapKey1).Return(nil)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.deleteAnswer(w, req)
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "get answer map, bad request",
			args: args{
				method: http.MethodGet,
				url:    answerPath,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.getAnswer(w, req)
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "get answer map, not found",
			args: args{
				method: http.MethodGet,
				url:    answerPath + answerKeyQueryString1,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Get(ctx, answerMapKey1).Return(nil, service.ErrResourceNotFound)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.getAnswer(w, req)
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name: "get answer map, ok",
			args: args{
				method: http.MethodGet,
				url:    answerPath + answerKeyQueryString1,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().Get(ctx, answerMapKey1).Return(answerMap1, nil)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.getAnswer(w, req)
			},
			wantStatus: http.StatusOK,
			wantBody:   answerMapString1,
		},
		{
			name: "get answer events, bad request",
			args: args{
				method: http.MethodGet,
				url:    answerEventPath,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.getAnswerEvents(w, req)
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "get answer events, not found",
			args: args{
				method: http.MethodGet,
				url:    answerEventPath + answerEventQueryString1,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().GetHistory(ctx, getHistoryRequest1).Return(nil, service.ErrResourceNotFound)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.getAnswerEvents(w, req)
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name: "get answer events, ok",
			args: args{
				method: http.MethodGet,
				url:    answerEventPath + answerEventQueryString1,
				body:   nil,
			},
			expectations: func(ctx context.Context, mock *mock_service.MockAnswerService) {
				mock.EXPECT().GetHistory(ctx, getHistoryRequest1).Return(answerEventList1, nil)
			},
			handler: func(app *App, w http.ResponseWriter, req *http.Request) {
				app.getAnswerEvents(w, req)
			},
			wantStatus: http.StatusOK,
			wantBody:   answerEventListString1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, tt.args.body)
			if err != nil {
				t.Fatalf("request init error = %v", err)
			}

			rr := httptest.NewRecorder()

			observerCore, _ := observer.New(zapcore.DebugLevel)

			db, _, err := sqlmock.New()
			if err != nil {
				t.Fatalf("db init error = %v", err)
			}
			defer db.Close()

			ctrl := gomock.NewController(t)
			service := mock_service.NewMockAnswerService(ctrl)

			if tt.expectations == nil {
				t.Fatalf("expectations not set")
			}
			tt.expectations(req.Context(), service)

			app := App{
				logger:  zap.New(observerCore),
				db:      db,
				service: service,
				router:  mux.NewRouter(),
			}
			app.registerRoutes()

			if tt.handler == nil {
				t.Fatalf("handler not set")
			}
			tt.handler(&app, rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("wrong status code, got = %d, wantStatus = %d", status, tt.wantStatus)
			}
			if body := rr.Body.String(); body != tt.wantBody {
				t.Errorf("wrong body, got = %s, wantBody = %s", body, tt.wantBody)
			}
		})
	}
}
