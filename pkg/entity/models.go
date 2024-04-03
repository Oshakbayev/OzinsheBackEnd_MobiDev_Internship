package entity

import (
	"errors"
	"time"
)

var (
	EmtpyString         = ""
	VerificationLinkURL = "http://localhost/auth/verifyAccount?link="
	AlreadyExist        = errors.New("user with this email already exist")
	DidNotFind          = errors.New("no rows in result set")
	ExpiredLink         = errors.New("link expired")
)

type User struct {
	Id              int
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	IsEmailVerified bool
}

type VerificationEmail struct {
	Id           int
	UserId       int
	EmailMessage string
	SecretCode   string
	IsEmailUsed  bool
	ExpTime      time.Time
}

type ErrorJSONResponse struct {
	Message string
}
