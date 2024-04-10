package service

import (
	"github.com/golang-jwt/jwt"
	"ozinshe/pkg/entity"
	"time"
)

func (s *Service) TokenGenerator(userID int, email string, role string) (string, error) {
	expTime := time.Now().Add(time.Hour * 48)
	claims := &entity.Claims{
		Email: email,
		Role:  role,
		Sub:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(entity.JWTKey)
	if err != nil {
		s.log.Printf("Error while signing jwt token: %v", err)
		return entity.EmtpyString, err
	}
	return signedToken, nil
}

func (s *Service) TokenChecker(tokenStr string) (*entity.Claims, error) {
	claims := &entity.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {

		return entity.JWTKey, nil
	})
	if err != nil {
		if err.Error() == jwt.ErrSignatureInvalid.Error() {
			s.log.Printf("Error in TokenChecker(Service): %v", err)
			return claims, err
		}
		s.log.Printf("Error in TokenChecker(Service): %v", err)
		return claims, err
	}
	if !tkn.Valid {
		s.log.Printf("Error in TokenChecker(Service): %v", err)
		return claims, err
	}
	decodedClaims := tkn.Claims.(*entity.Claims)
	return decodedClaims, nil
}
