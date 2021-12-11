package repository

import "books/domain"

type UserRepository interface {
	GetUser(username string) (domain.User, error)
}