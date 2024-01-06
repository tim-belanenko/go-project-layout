package exterrors

import (
	"errors"
	"fmt"
)

const (
	FailedParseLogLevelMsg        = "failed to parse log level"
	FailedGetLoggerFormContextMsg = "faile to get logger from context"
	FailedLoadMigrationFSMsg      = "failed to load migration fs"
	FailedMigrationMsg            = "failed to migrate"
)

var (
	ErrFailedParseLogLevel        = errors.New(FailedParseLogLevelMsg)
	ErrFailedGetLoggerFormContext = errors.New(FailedGetLoggerFormContextMsg)
	ErrFailedLoadMigrationsFS     = errors.New(FailedLoadMigrationFSMsg)
	ErrFailedMigration            = errors.New(FailedMigrationMsg)
)

func Error(operation string, errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	return fmt.Errorf("%s: %w", operation, errors.Join(errs...))
}
