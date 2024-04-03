package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"ozinshe/pkg/entity"
	"time"
)

type AuthService interface {
	SignUp(*entity.User) error
	VerifyAccount(string) error
}

func (s *Service) SignUp(user *entity.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("error while creating hash of password in SignUp(Service): %s", err.Error())
		return err
	}
	user.Password = string(hashedPass)
	ExistedUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		if err.Error() != "no rows in result set" {
			return err
		}
		if err = s.repo.CreateUser(user); err != nil {
			return err
		}
	} else if !ExistedUser.IsEmailVerified {
		// if user registered but not verified and try to register again
		user.Id = ExistedUser.Id
		if err = s.repo.UpdateUserByID(user); err != nil {
			return err
		}
	} else {
		return errors.New("user with this email already exist")
	}
	emailContent, secretCode, err := s.VerificationEmailGenerator(user.Email)
	if err != nil {
		s.log.Printf("error during verification email creation in SignUp(Service):", err.Error())
		return err
	}
	if err = s.SendVerificationEmail(user.Email, emailContent); err != nil {
		s.log.Printf("error during verification email sending in SignUp(Service):", err.Error())
		return err
	}
	if err = s.CreateVerificationEmail(user.Id, secretCode); err != nil {
		s.log.Printf("error during verification email creating in SignUp(Service):", err.Error())
		return err
	}
	return nil
}

func (s *Service) VerifyAccount(secretCode string) error {
	verificationEmail, err := s.repo.GetVerificationEmailStatusBySecretCode(secretCode)
	if err != nil {
		s.log.Printf("error during getting verification email in VerifyAccount(Service):", err.Error())
		return err
	}
	if verificationEmail.ExpTime.Before(time.Now()) {
		s.log.Printf("error in VerifyAccount(Service) %s", err.Error())
		err = s.repo.DeleteVerificationEmailByUserId(verificationEmail.UserId)
		if err != nil {
			s.log.Printf("error during deleting verification email in VerifyAccount(Service):", err.Error())
			return err
		}
		return errors.New("link expired")
	}
	err = s.repo.UpdateUsersEmailStatus(verificationEmail.UserId)
	if err != nil {
		s.log.Printf("error during updating user email status in VerifyAccount(Service):", err.Error())
		return err
	}
	err = s.repo.DeleteVerificationEmailByUserId(verificationEmail.UserId)
	if err != nil {
		s.log.Printf("error during deleting verification email in VerifyAccount(Service):", err.Error())
		return err
	}
	return nil
}
