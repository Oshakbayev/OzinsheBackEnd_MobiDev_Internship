package repository

import (
	"context"
	"ozinshe/pkg/entity"
)

type UserRepo interface {
	GetUserByEmail(string) (*entity.User, error)
	UpdateUserByID(*entity.User) error
	UpdateUsersEmailStatus(int) error
}

func (r *RepoStruct) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password, &user.IsEmailVerified)
	if err != nil {
		r.log.Printf("error in GetUserByEmail(repository):%s", err.Error())
	}
	return &user, err
}

func (r *RepoStruct) UpdateUserByID(user *entity.User) error {
	_, err := r.db.Exec(context.Background(), "UPDATE users SET password_hash = $1", user.Password)
	if err != nil {
		r.log.Printf("error in UpdateUserByID(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) UpdateUsersEmailStatus(userId int) error {
	_, err := r.db.Exec(context.Background(), "UPDATE users SET is_email_verified = true WHERE id=$1 ", userId)
	if err != nil {
		r.log.Printf("error in UpdateUsersEmailStatus(repository):%s", err.Error())
	}
	return err
}
