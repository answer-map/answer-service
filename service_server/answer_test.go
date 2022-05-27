package service_server_test

import (
	"context"
	"reflect"
	"testing"

	answer_protobuf "github.com/answer-map/answer-protobuf"
	"github.com/answer-map/answer-service/service_server"
	"github.com/volatiletech/null/v8"

	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

var (
	createReq1 = &answer_protobuf.CreateRequest{Answer: &answer_protobuf.Answer{Key: "name-o", Value: "bingo"}}
	createRes1 = &answer_protobuf.CreateResponse{StatusCode: 0}
)

func Test_answerServiceServer_Create(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *answer_protobuf.CreateRequest
	}
	tests := []struct {
		name         string
		args         args
		expectations func(mock sqlmock.Sqlmock)
		want         *answer_protobuf.CreateResponse
		wantErr      bool
	}{
		{
			name: "",
			args: args{
				in0: context.Background(),
				in1: createReq1,
			},
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("select * from public.answer_map where answer_map.key = ?").WithArgs(createReq1.Answer.Key).WillReturnRows(sqlmock.NewRows([]string{"key", "value"}).
					AddRow("name-o", null.NewString("", false)))
				mock.ExpectClose()
			},
			want: createRes1,
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

			server := service_server.NewAnswerServiceServer(db)

			got, err := server.Create(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("answerServiceServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerServiceServer.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
