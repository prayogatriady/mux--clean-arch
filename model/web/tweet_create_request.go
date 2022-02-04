package web

type TweetCreateRequest struct {
	Tweet string `json:"tweet" validate:"required,min=1,max=500"`
}
