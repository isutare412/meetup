package config

import (
	"fmt"

	glog "gorm.io/gorm/logger"
)

type LogFormat string

const (
	LogFormatJson LogFormat = "json"
	LogFormatText LogFormat = "text"
)

type PostgresLogLevel string

const (
	PostgresLogLevelSilent = "silent"
	PostgresLogLevelError  = "error"
	PostgresLogLevelWarn   = "warn"
	PostgresLogLevelInfo   = "info"
)

func (l PostgresLogLevel) Validate() error {
	switch l {
	case PostgresLogLevelSilent:
	case PostgresLogLevelError:
	case PostgresLogLevelWarn:
	case PostgresLogLevelInfo:
	default:
		return fmt.Errorf("unknown postgresLogLevel(%s)", l)
	}
	return nil
}

func (l PostgresLogLevel) IntoGORMLogLevel() glog.LogLevel {
	switch l {
	case PostgresLogLevelSilent:
		return glog.Silent
	case PostgresLogLevelError:
		return glog.Error
	case PostgresLogLevelWarn:
		return glog.Warn
	case PostgresLogLevelInfo:
		return glog.Info
	}
	return glog.Warn
}
