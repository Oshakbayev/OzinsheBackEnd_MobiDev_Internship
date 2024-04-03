package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/smtp"
	"ozinshe/pkg/entity"
	"time"
)

func (s *Service) VerificationEmailGenerator(email string) (string, string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	allowedChars := "0123456789"
	var result string
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(allowedChars))
		result += string(allowedChars[randomIndex])
	}
	for {
		secretCode, err := bcrypt.GenerateFromPassword([]byte(result+email), bcrypt.DefaultCost)
		if err != nil {
			return entity.EmtpyString, entity.EmtpyString, fmt.Errorf("error while hashing secretCode for email verification: error: %s", err)
		}
		_, err = s.repo.GetVerificationEmailStatusBySecretCode(string(secretCode))
		if err != nil {
			if err.Error() != "no rows in result set" {
				return entity.EmtpyString, entity.EmtpyString, fmt.Errorf("error during GetVerificationEmailStatusBySecretCode in VerificationEmailGenerator(Service): %s", err.Error())
			}
			emailContent := "Hello, thanks for registration on our app! Please follow the link attached below to complete your registration " + entity.VerificationLinkURL + string(secretCode) + " Reminder: the link is valid for 2 days"
			return emailContent, string(secretCode), nil
		}
	}
}

func (s *Service) SendVerificationEmail(emailAddress, emailContent string) error {
	to := []string{emailAddress}
	subject := "Subject: Ozinshe registration link\r\n"
	auth := smtp.PlainAuth("", "rakatan228322@gmail.com", "zgjw nlyp zyhk bczp", "smtp.gmail.com")
	msg := []byte(subject + "\r\n" + emailContent)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "rakatan228322@gmail.com", to, msg)
	if err != nil {
		return fmt.Errorf("error while sending verification email: error: %s", err)
	}
	return nil
}

func (s *Service) CreateVerificationEmail(userID int, verificationLink string) error {
	expTime := time.Now().Add(time.Hour * 48)
	email := entity.VerificationEmail{
		UserId:     userID,
		SecretCode: verificationLink,
		ExpTime:    expTime,
	}
	err := s.repo.CreateVerificationEmail(&email)
	if err != nil {
		return err
	}
	return nil
}
