package repository

import (
	"context"
	"ozinshe/pkg/entity"
)

type VerificationEmailRepo interface {
	CreateVerificationEmail(*entity.VerificationEmail) error
	GetVerificationEmailStatusBySecretCode(string) (entity.VerificationEmail, error)
	DeleteVerificationEmailByUserId(int) error
}

func (r *RepoStruct) CreateVerificationEmail(email *entity.VerificationEmail) error {
	_, err := r.db.Exec(context.Background(), "INSERT INTO verification_emails (user_id, secret_code, expired_at) VALUES ($1,$2,$3)", email.UserId, email.SecretCode, email.ExpTime)
	if err != nil {
		r.log.Printf("error during insert in CreateVerificationEmail(Repository): %s", err.Error())
	}
	return err
}

func (r *RepoStruct) UpdateVerificationEmailStatusBySecretCode(secretCode string) error {
	_, err := r.db.Exec(context.Background(), "UPDATE verification_emails SET is_used = true WHERE secret_code = $1", secretCode)
	if err != nil {
		r.log.Printf("error during update in UpdateVerificationEmailStatusBySecretCode(Repository): %s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetVerificationEmailStatusBySecretCode(secretCode string) (entity.VerificationEmail, error) {
	var email entity.VerificationEmail
	err := r.db.QueryRow(context.Background(), "SELECT * FROM verification_emails WHERE secret_code = $1", secretCode).Scan(&email.Id, &email.UserId, &email.SecretCode, &email.IsEmailUsed, &email.ExpTime)
	if err != nil {
		r.log.Printf("error during select in GetVerificationEmailStatusBySecretCode(Repository): %s", err.Error())
	}
	return email, err
}

func (r *RepoStruct) DeleteVerificationEmailByUserId(userId int) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM verification_emails WHERE user_id = $1", userId)
	if err != nil {
		r.log.Printf("error during update in UpdateVerificationEmailStatusBySecretCode(Repository): %s", err.Error())
	}
	return err
}
