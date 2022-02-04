package service

import (
	"context"
	"go-rest-api/model/web"
)

type TweetService interface {
	CreateTweet(ctx context.Context, request web.TweetCreateRequest) web.TweetResponse
	UpdateTweet(ctx context.Context, request web.TweetUpdateRequest) web.TweetResponse
	DeleteTweet(ctx context.Context, tweet_id int, username string) web.TweetResponse
	Home(ctx context.Context, username string) web.TweetResponse
}
