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
	GetUserEmailById(int) (string, error)
	ChangePasswordByEmail(email string, newPasswordHash string) error
}

func (r *RepoStruct) CreateUser(user *entity.User) error {
	err := r.db.QueryRow(context.Background(), "INSERT INTO users (email,password_hash,role,username,birth_date,phone_num) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id", user.Email, user.Password, "user", user.Username, user.BirthDate, user.PhoneNumber).Scan(&user.Id)
	if err != nil {
		r.log.Printf("error in CreateUser(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password, &user.IsEmailVerified, &user.Role, &user.Username, &user.BirthDate, &user.PhoneNumber)
	if err != nil {
		r.log.Printf("error in GetUserByEmail(repository):%s", err.Error())
	}
	return &user, err
}

func (r *RepoStruct) UpdateUserByID(user *entity.User) error {
	_, err := r.db.Exec(context.Background(), "UPDATE users SET username = $1,birth_date = $2,phone_num = $3", user.Username, user.BirthDate, user.PhoneNumber)
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

func (r *RepoStruct) ChangePasswordByEmail(email string, newPasswordHash string) error {
	_, err := r.db.Exec(context.Background(), "UPDATE users SET password_hash = $1 WHERE email=$2 ", newPasswordHash, email)
	if err != nil {
		r.log.Printf("error in ChangePasswordByUserId(repository):%s", err.Error())
	}
	return err
}
func (r *RepoStruct) GetUserEmailById(userId int) (string, error) {
	query := `SELECT email FROM users WHERE id =$1`
	var email string
	err := r.db.QueryRow(context.Background(), query, userId).Scan(&email)
	if err != nil {
		r.log.Printf("error in GetUserEmailById(repository):%s", err.Error())
	}
	return email, err
}
