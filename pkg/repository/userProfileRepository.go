package repository

import (
	"context"
	"ozinshe/pkg/entity"
)

type UserProfileRepo interface {
	GetUserProfileByUserId(int) (entity.UserProfile, error)
	CreateUserProfile(*entity.UserProfile) error
	UpdateUserProfile(*entity.UserProfile) error
}

func (r *RepoStruct) GetUserProfileByUserId(userId int) (entity.UserProfile, error) {
	query := `SELECT * FROM user_profile WHERE user_id = $1`
	var userProfile entity.UserProfile
	err := r.db.QueryRow(context.Background(), query, userId).Scan(&userProfile.Id, &userProfile.BirthDate, &userProfile.UserId, &userProfile.Language, &userProfile.PhoneNumber)
	if err != nil {
		r.log.Printf("error in GetUserProfileByUserId(repository):%s", err.Error())
	}
	return userProfile, err
}

func (r *RepoStruct) CreateUserProfile(userProfile *entity.UserProfile) error {
	query := `INSERT INTO user_profile (birth_date,user_id,language,phone_number) VALUES ($1,$2,$3,$4)`
	_, err := r.db.Exec(context.Background(), query, userProfile.BirthDate, userProfile.UserId, userProfile.Language, userProfile.PhoneNumber)
	if err != nil {
		r.log.Printf("error in CreateUserProfile(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) UpdateUserProfile(userProfile *entity.UserProfile) error {
	query := `UPDATE user_profile SET birth_date=$1,language = $2,phone_number = $3 WHERE user_id  = $4`
	_, err := r.db.Exec(context.Background(), query, userProfile.BirthDate, userProfile.Language, userProfile.PhoneNumber, userProfile.UserId)
	if err != nil {
		r.log.Printf("error in UpdateUserProfile(repository):%s", err.Error())
	}
	return err
}
