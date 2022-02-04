package web

type TweetResponse struct {
	TweetId  int    `json:"tweet_id"`
	Username string `json:"username"`
	Tweet    string `json:"tweet"`
}
