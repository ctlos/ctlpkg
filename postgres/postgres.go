package postgres

import (
	"context"
	"errors"

	"gitea.dancheg97.ru/dancheg97/go-sqlc/gen/sqlc"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgres struct {
	*pgxpool.Pool
	*sqlc.Queries
}

func Get(conn string) (*Postgres, error) {
	config, err := pgxpool.ParseConfig(conn)
	if err != nil {
		return nil, err
	}

	config.ConnConfig.LogLevel = pgx.LogLevelError
	config.ConnConfig.Logger = logrusadapter.NewLogger(logrus.StandardLogger())
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	sqlc := sqlc.New(pool)
	return &Postgres{
		Pool:    pool,
		Queries: sqlc,
	}, nil
}

func (pg *Postgres) WrapTx(ctx context.Context, txFunc func(p *Postgres) error) error {
	tx, err := pg.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	dbtx := pg.Queries.WithTx(tx)
	defer func() {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			if !errors.Is(rollbackErr, pgx.ErrTxClosed) {
				logrus.Error(`error rolling transaction back`, rollbackErr)
			}
		}
	}()
	err = txFunc(&Postgres{
		Pool:    pg.Pool,
		Queries: dbtx,
	})
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (pg *Postgres) CheckConstraint(err error, cst string) bool {
	if err == nil {
		return false
	}
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if ok {
		return pgErr.ConstraintName == cst
	}
	return false
}
