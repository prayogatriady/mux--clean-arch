package helper

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Username:  user.Username,
		Password:  user.Password,
		GroupUser: user.GroupUser,
		Email:     user.Email,
	}
}
