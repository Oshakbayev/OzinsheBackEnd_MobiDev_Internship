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
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Movie struct {
	Id               int
	CategoryIDs      []int
	CategoryAgeIDs   []int
	CreatedDate      time.Time
	Description      string
	Director         string
	Favorite         bool
	GenreIDs         []int
	Keywords         string
	LastModifiedDate time.Time
	MovieType        string
	Name             string
	Poster           []Image
	Producer         string
	Screenshots      []Image
	SeasonCount      int
	SeriesCount      int
	Timing           int
	Trend            bool
	Video            Video
	WatchCount       int
	Year             int
}

type Genre struct {
	Id         int
	Name       string
	MovieCount int
}

type Category struct {
	MovieId    int
	Id         int
	Link       string
	MovieCount int
	Name       string
}

type Image struct {
	Id      int
	FileId  int
	MovieId int
	Link    string
}

type Video struct {
	Id           int
	Link         string
	SeriesNumber int
	SeasonId     int
}
