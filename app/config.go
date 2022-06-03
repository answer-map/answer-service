package app

import (
	"fmt"
	"os"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	HTTP           *HTTPConfig               `json:"http" yaml:"http"`
	ZapLogger      *zap.Config               `json:"zapLogger" yaml:"zapLogger"`
	AnswerDataBase *PostgreSQLDataBaseConfig `json:"answerDatabase" yaml:"answerDatabase"`
}

// ConfigFromEnvironmentVariables returns a valid Config from the enviroment, or an error.
func ConfigFromEnvironmentVariables() (*Config, error) {
	atomicLevel := zap.NewAtomicLevel()

	if os.Getenv("LOGGER_LEVEL") == "debug" {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	httpPort, err := strconv.ParseUint(os.Getenv("HTTP_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	dbPort, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    os.Getenv("LOGGER_ENCODER_CONFIG_MESSAGE_KEY"),
		LevelKey:      os.Getenv("LOGGER_ENCODER_CONFIG_LEVEL_KEY"),
		TimeKey:       os.Getenv("LOGGER_ENCODER_CONFIG_TIME_KEY"),
		NameKey:       os.Getenv("LOGGER_ENCODER_CONFIG_NAME_KEY"),
		CallerKey:     os.Getenv("LOGGER_ENCODER_CONFIG_CALLER_KEY"),
		FunctionKey:   os.Getenv("LOGGER_ENCODER_CONFIG_FUNCTION_KEY"),
		StacktraceKey: os.Getenv("LOGGER_ENCODER_CONFIG_STACKTRACE_KEY"),
	}

	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	zapLoggerConfig := zap.Config{
		Level:            atomicLevel,
		Development:      os.Getenv("LOGGER_DEVELOPMENT") == "true",
		Encoding:         os.Getenv("LOGGER_ENCODING"),
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{os.Getenv("LOGGER_STDOUT_PATH")},
		ErrorOutputPaths: []string{os.Getenv("LOGGER_STDERR_PATH")},
	}

	config := &Config{
		HTTP: &HTTPConfig{
			Port: httpPort,
		},
		ZapLogger: &zapLoggerConfig,
		AnswerDataBase: &PostgreSQLDataBaseConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			DBName:   os.Getenv("DATABASE_NAME"),
			SSLMode:  os.Getenv("DATABASE_SSL_MODE"),
		},
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
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
	Port uint64 `json:"port" yaml:"port"`
}

func (httpConfig *HTTPConfig) Validate() error {
	port := httpConfig.Port

	if port != uint64(80) && port != uint64(443) && port < uint64(1025) {
		return fmt.Errorf("config error: invalid http port")
	}

	return nil
}

func (httpConfig *HTTPConfig) Address() string {
	return fmt.Sprintf(":%d", httpConfig.Port)
}

type PostgreSQLDataBaseConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     uint64 `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"dbName" yaml:"dbName"`
	SSLMode  string `json:"sslMode" yaml:"sslMode"`
}

func (dbConfig *PostgreSQLDataBaseConfig) Validate() error {
	if err := validation.ValidateStruct(dbConfig,
		validation.Field(&dbConfig.Host, validation.Required),
		validation.Field(&dbConfig.Port, validation.Required, validation.Min(uint64(1024))),
		validation.Field(&dbConfig.User, validation.Required),
		validation.Field(&dbConfig.Password, validation.Required),
		validation.Field(&dbConfig.DBName, validation.Required),
		validation.Field(&dbConfig.SSLMode, validation.Required, validation.In("disable"))); err != nil {
		return err
	}

	return nil
}

func (dbConfig *PostgreSQLDataBaseConfig) DataSource() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)
}
