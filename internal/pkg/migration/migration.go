package migration

import (
	"embed"
	"errors"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"layout/internal/pkg/exterrors"
)

//go:embed postgres/*.sql
var fs embed.FS

type Migration struct {
	migrate *migrate.Migrate
}

func New(pgDSN string) (*Migration, error) {
	const op = "migration.New"

	driver, err := iofs.New(fs, "postgres")
	if err != nil {
		return nil, exterrors.Error(op, exterrors.ErrFailedMigration, err)
	}

	migrage, err := migrate.NewWithSourceInstance(
		"iofs",
		driver,
		pgDSN,
	)
	if err != nil {
		return nil, exterrors.Error(op, exterrors.ErrFailedMigration, err)
	}

	return &Migration{
		migrate: migrage,
	}, nil
}

func (m *Migration) Up() error {
	const op = "migration.Up"

	if err := m.migrate.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return exterrors.Error(op, err, exterrors.ErrFailedMigration)
	}

	return nil
}

func (m *Migration) Down() error {
	const op = "migration.Down"
	if err := m.migrate.Down(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return exterrors.Error(op, err, exterrors.ErrFailedMigration, err)
	}
	return nil
}
