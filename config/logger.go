package config

import "go.uber.org/zap"

type LoggerClient struct {
	*zap.Logger
}
