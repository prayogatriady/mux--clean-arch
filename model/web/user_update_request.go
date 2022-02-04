package web

type UserUpdateRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password" validate:"required,min=4,max=20"`
	GroupUser string `json:"group_user" validate:"required,min=1,max=30"`
	Email     string `json:"email" validate:"email,min=1,max=100"`
}
