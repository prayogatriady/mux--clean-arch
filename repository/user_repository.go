package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindByUsername(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByUsernamePassword(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}
