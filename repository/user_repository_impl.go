package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepo *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "insert into user (username, password, group_user, email) VALUES (?,?,?,?)"
	_, err := tx.ExecContext(ctx, query, user.Username, user.Password, user.GroupUser, user.Email)
	helper.PanicIfErr(err)

	return user
}

func (userRepo *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "update user set password = ?, group_user = ?, email = ? where username = ?"
	_, err := tx.ExecContext(ctx, query, user.Password, user.GroupUser, user.Email, user.Username)
	helper.PanicIfErr(err)

	return user
}

func (userRepo *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	query := "delete from user where username = ?"
	_, err := tx.ExecContext(ctx, query, user.Username)
	helper.PanicIfErr(err)
}

func (userRepo *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "select username, password, group_user, email from user where username = ?"
	rows, err := tx.QueryContext(ctx, query, user.Username)
	helper.PanicIfErr(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Username, &user.Password, &user.GroupUser, &user.Email)
		helper.PanicIfErr(err)
		return user, nil
	} else {
		user.Username = ""
		return user, errors.New("User is not found")
	}
}

func (userRepo *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "select username, password, group_user, email from user"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfErr(err)
	defer rows.Close()

	users := []domain.User{}
	user := domain.User{}
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Password, &user.GroupUser, &user.Email)
		helper.PanicIfErr(err)

		users = append(users, user)
	}

	return users
}
