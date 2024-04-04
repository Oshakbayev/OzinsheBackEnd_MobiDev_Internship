package entity

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	EmtpyString         = ""
	VerificationLinkURL = "http://localhost/auth/verifyAccount?link="
	AlreadyExist        = "499" //"user with this email already exist")
	DidNotFind          = "no rows in result set"
	ExpiredLink         = "497" //"link expired"
	NotVerifiedEmail    = "496"
	InvalidPassword     = "495"
	InvalidEmail        = "494"
	JWTKey              = []byte("sercet_key")
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

type Credentials struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Claims struct {
	Email string `json:"email"`
	Sub   int    `json:"sub"`
	Level string `json:"level"`
	jwt.StandardClaims
}
