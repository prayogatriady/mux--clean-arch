package web

type TweetUpdateRequest struct {
	Username string `json:"username" validate:"required,min=4,max=30"`
	Tweet    string `json:"tweet" validate:"required,min=1,max=500"`
}
