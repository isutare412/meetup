package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	glog "gorm.io/gorm/logger"
)

type GORM struct {
	Level         glog.LogLevel
	SlowThreshold time.Duration
}

func NewGORM(level glog.LogLevel) *GORM {
	return &GORM{
		Level:         level,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l *GORM) LogMode(level glog.LogLevel) glog.Interface {
	newLogger := *l
	newLogger.Level = level
	return &newLogger
}

func (l *GORM) Info(_ context.Context, msg string, args ...interface{}) {
	if l.Level >= glog.Info {
		l.base().Infof(msg, args...)
	}
}

func (l *GORM) Warn(_ context.Context, msg string, args ...interface{}) {
	if l.Level >= glog.Warn {
		l.base().Warnf(msg, args...)
	}
}

func (l *GORM) Error(_ context.Context, msg string, args ...interface{}) {
	if l.Level >= glog.Error {
		l.base().Errorf(msg, args...)
	}
}

func (l *GORM) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.Level <= glog.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.Level >= glog.Error && !errors.Is(err, glog.ErrRecordNotFound):
		sql, rows := fc()

		log := l.base().With(
			"queryMS", fmt.Sprintf("%.3f", float64(elapsed.Nanoseconds())/1e6),
			"sql", sql,
		)
		if rows == -1 {
			log = log.With("rows", "-")
		} else {
			log = log.With("rows", rows)
		}
		log.Error(err)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.Level >= glog.Warn:
		sql, rows := fc()

		log := l.base().With(
			"queryMS", fmt.Sprintf("%.3f", float64(elapsed.Nanoseconds())/1e6),
			"sql", sql,
		)
		if rows == -1 {
			log = log.With("rows", "-")
		} else {
			log = log.With("rows", rows)
		}
		log.Warn("SLOW SQL >= %v", l.SlowThreshold)
	case l.Level == glog.Info:
		sql, rows := fc()

		log := l.base().With(
			"queryMS", fmt.Sprintf("%.3f", float64(elapsed.Nanoseconds())/1e6),
			"sql", sql,
		)
		if rows == -1 {
			log = log.With("rows", "-")
		} else {
			log = log.With("rows", rows)
		}
		log.Info("query trace")
	}
}

func (l *GORM) base() *zap.SugaredLogger {
	return S().WithOptions(zap.AddCallerSkip(2))
}
