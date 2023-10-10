package config

import "go.uber.org/zap"

type LoggerClient struct {
	*zap.Logger
}

// https://github.com/uber-go/zap
func NewLoggerClient() (*LoggerClient, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &LoggerClient{logger}, nil
}
