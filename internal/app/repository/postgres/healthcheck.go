package postgres

import (
	"context"
	"layout/internal/app/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *Postgres {
	return &Postgres{
		conn: conn,
	}
}

func (p *Postgres) Healthcheck(ctx context.Context) (*usecase.Module, error) {
	m := &usecase.Module{
		Module: "postgres",
		Status: usecase.StatusInactive,
	}

	err := p.conn.Ping(ctx)
	if err != nil {
		return m, err
	}

	m.Status = usecase.StatusActive
	return m, nil
}
