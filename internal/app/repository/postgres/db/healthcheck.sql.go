// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: healthcheck.sql

package db

import (
	"context"
)

const healthcheck = `-- name: Healthcheck :exec
SELECT 1
`

func (q *Queries) Healthcheck(ctx context.Context) error {
	_, err := q.db.Exec(ctx, healthcheck)
	return err
}