package service

import (
	"books/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtService struct {
	secret     string
}

func NewJwtService() *JwtService{
	return &JwtService{
		secret: "secret",
	}
}

func (s *JwtService) GeneratoToken(user domain.User) (string, error ){
	claims := jwt.RegisteredClaims{
		Issuer: "books-server",
		Subject: user.Username,
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.secret))

	if err != nil {
		return "", errors.New("could not generate token: " + err.Error())
	}

	return signedToken, nil
}

func (s *JwtService) ValidateToken(tokenString string) (*jwt.Token, error){
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		
		}
		return []byte(s.secret), nil
	});
}