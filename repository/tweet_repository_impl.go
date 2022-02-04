package repository

import (
	"context"
	"database/sql"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type TweetRepositoryImpl struct {
}

func (tweetRepo *TweetRepositoryImpl) Post(ctx context.Context, tx *sql.Tx, tweet domain.Tweet) domain.Tweet {
	query := "insert into tweet (username, tweet) VALUES (?,?,?)"
	result, err := tx.ExecContext(ctx, query, tweet.Username, tweet.Tweet)
	helper.PanicIfErr(err)

	tweet_id, err := result.LastInsertId()
	tweet.TweetId = int(tweet_id)

	return tweet
}

func (tweetRepo *TweetRepositoryImpl) Edit(ctx context.Context, tx *sql.Tx, tweet domain.Tweet) domain.Tweet {
	query := "update tweet set tweet = ? where tweet_id = ? and username = ?"
	_, err := tx.ExecContext(ctx, query, tweet.Tweet, tweet.TweetId, tweet.Username)
	helper.PanicIfErr(err)

	return tweet
}

func (tweetRepo *TweetRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, tweet domain.Tweet) {
	query := "delete from tweet where tweet_id = ? and username = ?"
	_, err := tx.ExecContext(ctx, query, tweet.TweetId, tweet.Username)
	helper.PanicIfErr(err)
}

func (tweetRepo *TweetRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, tweet domain.Tweet) []domain.Tweet {
	query := "select tweet_id, username, tweet from tweet where username = ?"
	rows, err := tx.QueryContext(ctx, query, tweet.Username)
	helper.PanicIfErr(err)
	defer rows.Close()

	tweets := []domain.Tweet{}
	for rows.Next() {
		err := rows.Scan(&tweet.TweetId, &tweet.Username, &tweet.Tweet)
		helper.PanicIfErr(err)

		tweets = append(tweets, tweet)
	}

	return tweets
}
