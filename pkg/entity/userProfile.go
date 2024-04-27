package entity

type UserProfile struct {
	Id          int
	BirthDate   string `json:"birthDate"`
	UserId      int
	Language    string `json:"language"`
	PhoneNumber string `json:"phoneNumber"`
}

type User struct {
	Id              int
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	IsEmailVerified bool
	Role            string
	BirthDate       string `json:"birthDate"`
	PhoneNumber     string `json:"phoneNumber"`
}
