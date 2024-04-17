package entity

type UserProfile struct {
	Id          int
	BirthDate   string `json:"birthDate"`
	UserId      int
	Language    string `json:"language"`
	PhoneNumber string `json:"phoneNumber"`
}
