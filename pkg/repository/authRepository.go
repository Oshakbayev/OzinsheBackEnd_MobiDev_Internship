package repository

import (
	"context"
	"ozinshe/pkg/entity"
)

type AuthRepo interface {
	CreateUser(*entity.User) error
}

func (r *RepoStruct) CreateUser(user *entity.User) error {
	err := r.db.QueryRow(context.Background(), "INSERT INTO users (email,password_hash) VALUES ($1,$2) RETURNING id", user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		r.log.Printf("error in CreateUser(repository):%s", err.Error())
	}
	return err
}
