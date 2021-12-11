package service

import (
	"books/domain"
	"books/repository"
	"errors"
)

type AuthenticationService struct {
	users repository.UserRepository
}

func NewAuthenticationService(users repository.UserRepository) *AuthenticationService{
	return &AuthenticationService{
		users: users,
	}
}

func (auth *AuthenticationService) AuthenticateUser(username string, pswd string) (domain.User, error) {
	user, err := auth.users.GetUser(username)

	if err != nil {
		return user, errors.New("could not authenticate " + user.Username + ": " + err.Error())
	}

	if user.Password != pswd {
		return user, errors.New("invalid credentials")
	}

	return user, nil
} 