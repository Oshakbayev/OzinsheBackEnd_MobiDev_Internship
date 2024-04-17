package repository

import (
	"context"
	"ozinshe/pkg/entity"
)

type UserRepo interface {
	CreateUser(*entity.User) error
	GetUserByEmail(string) (*entity.User, error)
	UpdateUserByID(*entity.User) error
	UpdateUsersEmailStatus(int) error
	GetPasswordByUserId(int) (string, error)
	ChangePasswordByUserId(int, string) error
}

func (r *RepoStruct) CreateUser(user *entity.User) error {
	err := r.db.QueryRow(context.Background(), "INSERT INTO users (email,password_hash,role) VALUES ($1,$2,$3) RETURNING id", user.Email, user.Password, "user").Scan(&user.Id)
	if err != nil {
		r.log.Printf("error in CreateUser(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password, &user.IsEmailVerified, &user.Role)
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

func (r *RepoStruct) GetPasswordByUserId(userId int) (string, error) {
	var passwordHash string
	err := r.db.QueryRow(context.Background(), "SELECT password_hash FROM users WHERE id=$1 ", userId).Scan(&passwordHash)
	if err != nil {
		r.log.Printf("error in ChangePasswordByUserId(repository):%s", err.Error())
	}
	return passwordHash, err
}

func (r *RepoStruct) ChangePasswordByUserId(userId int, newPasswordHash string) error {
	_, err := r.db.Exec(context.Background(), "UPDATE users SET password_hash = $1 WHERE id=$2 ", newPasswordHash, userId)
	if err != nil {
		r.log.Printf("error in ChangePasswordByUserId(repository):%s", err.Error())
	}
	return err
}
