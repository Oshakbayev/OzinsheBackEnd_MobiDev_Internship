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
	ErrNoRows           = "no rows in result set"
)

const (
	BucketName                   = "test-bucket-dostap"
	VerificationSecretCodeLength = 16
	Charset                      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	PicturesLinkNameLength       = 32
)

type User struct {
	Id              int
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	IsEmailVerified bool
	Role            string
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
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Image struct {
	Id      int
	FileId  int
	MovieId int
	Link    string
}
