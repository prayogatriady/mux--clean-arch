package web

type UserResponse struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	GroupUser string `json:"group_user"`
	Email     string `json:"email"`
}
