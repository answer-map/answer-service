package app_test

import (
	"testing"

	"github.com/answer-map/answer-service/app"
	"go.uber.org/zap"
)

var (
	devZapConfig = zap.NewDevelopmentConfig()
	dbConfig1    = &app.PostgreSQLDataBaseConfig{
		Host:     "::1",
		Port:     uint64(5432),
		User:     "lewis",
		Password: "open5esame",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
)

func TestPostgreSQLDataBaseConfig_Validate(t *testing.T) {
	tests := []struct {
		name     string
		dbConfig *app.PostgreSQLDataBaseConfig
		wantErr  bool
	}{
		{
			name:     "pass",
			dbConfig: dbConfig1,
			wantErr:  false,
		},
		{
			name: "invalid user",
			dbConfig: &app.PostgreSQLDataBaseConfig{
				User:    "",
				DBName:  "postgres",
				SSLMode: "disable",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dbConfig.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("PostgreSQLDataBaseConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgreSQLDataBaseConfig_DataSource(t *testing.T) {
	tests := []struct {
		name     string
		dbConfig *app.PostgreSQLDataBaseConfig
		want     string
	}{
		{
			name:     "pass",
			dbConfig: dbConfig1,
			want:     "host=::1 port=5432 user=lewis password=open5esame dbname=postgres sslmode=disable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dbConfig.DataSource(); got != tt.want {
				t.Errorf("PostgreSQLDataBaseConfig.DataSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPConfig_Validate(t *testing.T) {
	tests := []struct {
		name       string
		httpConfig *app.HTTPConfig
		wantErr    bool
	}{
		{
			name: "pass",
			httpConfig: &app.HTTPConfig{
				Port: 8080,
			},
			wantErr: false,
		},
		{
			name: "invalid port",
			httpConfig: &app.HTTPConfig{
				Port: 1024,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.httpConfig.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("HTTPConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTPConfig_Address(t *testing.T) {
	tests := []struct {
		name       string
		httpConfig *app.HTTPConfig
		want       string
	}{
		{
			name: "pass",
			httpConfig: &app.HTTPConfig{
				Port: 8080,
			},
			want: ":8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.httpConfig.Address(); got != tt.want {
				t.Errorf("HTTPConfig.Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *app.Config
		wantErr bool
	}{
		{
			name: "pass",
			config: &app.Config{
				HTTP:           &app.HTTPConfig{Port: 8080},
				AnswerDataBase: dbConfig1,
				ZapLogger:      &devZapConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.config.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
