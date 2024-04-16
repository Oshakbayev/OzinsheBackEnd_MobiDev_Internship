package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"ozinshe/pkg/entity"
	"regexp"
	"time"
)

type AuthService interface {
	SignUp(*entity.User) error
	VerifyAccount(string) error
	SigIn(*entity.Credentials) (*entity.User, error)
	TokenGenerator(int, string, string) (string, error)
}

func (s *Service) SignUp(user *entity.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("error while creating hash of password in SignUp(Service): %s", err.Error())
		return err
	}
	user.Password = string(hashedPass)
	regex := regexp.MustCompile(`.*[a-zA-Z0-9]+.*@.*\..*`)
	if !regex.MatchString(user.Email) {
		return errors.New("494")
	}
	ExistedUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		if err.Error() != entity.DidNotFind {
			return err
		}
		if err = s.repo.CreateUser(user); err != nil {
			return err
		}
	} else if !ExistedUser.IsEmailVerified {
		// if user registered but not verified and try to register again
		err = s.repo.DeleteVerificationEmailByUserId(ExistedUser.Id)
		if err != nil {
			s.log.Printf("error during deleting verification email in VerifyAccount(Service):", err.Error())
			return err
		}
		user.Id = ExistedUser.Id
		if err = s.repo.UpdateUserByID(user); err != nil {
			return err
		}
	} else {
		return errors.New(entity.AlreadyExist)
	}
	emailContent, secretCode, err := s.VerificationEmailGenerator(user.Email)
	if err != nil {
		s.log.Printf("error during verification email creation in SignUp(Service):", err.Error())
		return err
	}
	if err = s.SendVerificationEmail(user.Email, emailContent); err != nil {
		s.log.Printf("error during verification email sending in SignUp(Service):", err.Error())
		return errors.New("Invalid email")
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
		return errors.New("497")
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

func (s *Service) SigIn(credentials *entity.Credentials) (*entity.User, error) {
	user, err := s.repo.GetUserByEmail(credentials.Email)
	if err != nil {
		// can be now rows in result set
		return nil, err
	} else if !user.IsEmailVerified {
		//not verified email
		return nil, errors.New("email not verified")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		s.log.Printf("given password  is incorrect: %s", credentials.Password)
		return nil, fmt.Errorf("given password is incorrect: %s", credentials.Password)
	}
	return user, nil
}
