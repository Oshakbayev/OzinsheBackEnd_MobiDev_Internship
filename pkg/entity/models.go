package entity

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	EmtpyString            = ""
	VerificationLinkURL    = "http://localhost/auth/verifyAccount?link="
	AlreadyExist           = "user with this email already exist"
	DidNotFind             = "no rows in result set"
	ExpiredLink            = "this Link Expired"
	NotVerifiedEmail       = "this account's email not verified"
	InvalidPassword        = "invalid password"
	InvalidEmail           = "invalid email"
	InvalidConfirmPassword = "passwords must be same"
	JWTKey                 = []byte("sercet_key")
	ErrNoRows              = "no rows in result set"
)

const (
	VerificationSecretCodeLength = 16
	Charset                      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	UploadLinkNameLength         = 32
	UploadedFilesPath            = "assets/uploads/"
)

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
