package service

import (
	"context"
	"go-rest-api/model/web"
)

type UserService interface {
	CreateUser(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	UpdateUser(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	DeleteUser(ctx context.Context, username string)
	FindUser(ctx context.Context, username string) web.UserResponse
	FindAllUser(ctx context.Context) []web.UserResponse
	LoginUser(ctx context.Context, request web.UserLoginRequest) web.UserResponse
}
