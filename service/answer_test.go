package service_test

import (
	"context"
	"testing"

	"github.com/answer-map/answer-service/model"
	"github.com/answer-map/answer-service/service"

	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

var (
	createReq1 = &model.CreateRequest{AnswerKey: "name", AnswerValue: ""}
)

func Test_answerService_Create(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *model.CreateRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		wantErr      bool
	}{
		{
			name: "create name with empty value",
			args: args{
				in0: context.Background(),
				in1: createReq1,
			},
			expectations: func(mock sqlmock.Sqlmock) {},
			wantErr:      true,
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

			if err := server.Create(tt.args.in0, tt.args.in1); (err != nil) != tt.wantErr {
				t.Errorf("answerServiceServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
