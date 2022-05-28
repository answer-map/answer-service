package app

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
)

type Config struct {
	HTTP           *HTTPConfig               `json:"http" yaml:"http"`
	AnswerDataBase *PostgreSQLDataBaseConfig `json:"answerDatabase" yaml:"answerDatabase"`
	ZapLogger      *zap.Config               `json:"zapLogger" yaml:"zapLogger"`
}

func (config *Config) Validate() error {
	if err := validation.ValidateStruct(config,
		validation.Field(&config.HTTP, validation.Required),
		validation.Field(&config.AnswerDataBase, validation.Required),
		validation.Field(&config.ZapLogger, validation.Required)); err != nil {
		return fmt.Errorf("failed to find config fields, error = %v", err)
	}

	if err := config.HTTP.Validate(); err != nil {
		return fmt.Errorf("failed to validate http, error = %v", err)
	}

	if err := config.AnswerDataBase.Validate(); err != nil {
		return fmt.Errorf("failed to validate party data base, error = %v", err)
	}

	return nil
}

type HTTPConfig struct {
	Port uint16 `json:"port" yaml:"port"`
}

func (httpConfig *HTTPConfig) Validate() error {
	port := httpConfig.Port

	if port != 80 && port != 443 && port < 1025 {
		return fmt.Errorf("config error: invalid http port")
	}

	return nil
}

func (httpConfig *HTTPConfig) Address() string {
	return fmt.Sprintf(":%d", httpConfig.Port)
}

type PostgreSQLDataBaseConfig struct {
	User    string `json:"user" yaml:"user"`
	DBName  string `json:"dbName" yaml:"dbName"`
	SSLMode string `json:"sslMode" yaml:"sslMode"`
}

func (dbConfig *PostgreSQLDataBaseConfig) Validate() error {
	if err := validation.ValidateStruct(dbConfig,
		validation.Field(&dbConfig.User, validation.Required),
		validation.Field(&dbConfig.DBName, validation.Required),
		validation.Field(&dbConfig.SSLMode, validation.Required, validation.In("disable"))); err != nil {
		return err
	}

	return nil
}

func (dbConfig *PostgreSQLDataBaseConfig) DataSource() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbConfig.User, dbConfig.DBName, dbConfig.SSLMode)
}
