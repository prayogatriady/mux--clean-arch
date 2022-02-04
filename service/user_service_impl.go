package service

import (
	"context"
	"database/sql"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
	"go-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       &validate,
	}
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username:  request.Username,
		Password:  request.Password,
		GroupUser: request.GroupUser,
		Email:     request.Email,
	}

	user = service.UserRepository.Insert(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, domain.User{Username: request.Username})
	helper.PanicIfErr(err)

	user.Password = request.Password
	user.GroupUser = request.GroupUser
	user.Email = request.Email

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) DeleteUser(ctx context.Context, username string) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, domain.User{Username: username})
	helper.PanicIfErr(err)

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindUser(ctx context.Context, username string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, domain.User{Username: username})
	helper.PanicIfErr(err)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAllUser(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.ToUserResponse(user))
	}

	return userResponses
}

func (service *UserServiceImpl) LoginUser(ctx context.Context, request web.UserLoginRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err = service.UserRepository.FindByUsernamePassword(ctx, tx, user)
	helper.PanicIfErr(err)

	return helper.ToUserResponse(user)
}
