package repository

import "books/domain"

type BookRepository interface {
	GetAll() []domain.Book
	Save(b domain.Book) (domain.Book, error)
	Update(b domain.Book) (domain.Book, error)
	GetByISBN(isbn string) (domain.Book, error)
	GetById(id interface{}) (domain.Book, error)
}