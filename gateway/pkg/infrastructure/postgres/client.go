package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ctxKeyTx struct{}

var txKey ctxKeyTx

type Client struct {
	db *gorm.DB
}

func NewClient(ctx context.Context, cfg *config.PostgresConfig) (*Client, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &logger.GORM{
			Level:         cfg.Log.Level.IntoGORMLogLevel(),
			SlowThreshold: cfg.Log.SlowQuery,
		},
	})
	if err != nil {
		return nil, err
	}

	if sqlDB, err := db.DB(); err != nil {
		return nil, fmt.Errorf("getting *sql.DB from GORM: %w", err)
	} else if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("pinging DB: %w", err)
	}

	return &Client{
		db: db,
	}, nil
}

func (c *Client) MigrateSchema(ctx context.Context) error {
	success := make(chan struct{})
	fails := make(chan error, 1)
	go func() {
		err := c.db.AutoMigrate(
			&domain.User{},
		)
		if err != nil {
			fails <- err
			return
		}
		close(success)
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("time out while schema migration")
	case err := <-fails:
		return err
	case <-success:
	}
	return nil
}

func (c *Client) BeginTx(ctx context.Context, opts ...*sql.TxOptions) (
	ctxWithTx context.Context,
	commit, rollback func() error,
) {
	tx := c.db.WithContext(ctx).Begin(opts...)

	ctxWithTx = context.WithValue(ctx, txKey, tx)
	commit = func() error { return tx.Commit().Error }
	rollback = func() error { return tx.Rollback().Error }
	return ctxWithTx, commit, rollback
}

func (c *Client) extractTxOrDB(ctx context.Context) *gorm.DB {
	tx := ctx.Value(txKey)
	if tx == nil {
		return c.db
	}
	return tx.(*gorm.DB)
}
