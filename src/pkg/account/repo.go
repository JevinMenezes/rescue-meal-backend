package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

// ErrRepo - comment for export
var ErrRepo = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo - comment for export
func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `INSERT INTO users (email, password) VALUES ($1, $2)`
	if user.Email == "" || user.Password == "" {
		return ErrRepo
	}

	_, err := repo.db.ExecContext(ctx, sql, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
