package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type TweetRepository interface {
	Post(ctx context.Context, tx *sql.Tx, tweet domain.Tweet) domain.Tweet
	Edit(ctx context.Context, tx *sql.Tx, tweet domain.Tweet) domain.Tweet
	Delete(ctx context.Context, tx *sql.Tx, tweet domain.Tweet)
	FindAll(ctx context.Context, tx *sql.Tx, tweet domain.Tweet) []domain.Tweet
}
