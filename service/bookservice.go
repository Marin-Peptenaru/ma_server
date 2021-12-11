package service

import (
	"books/domain"
	"books/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func New(repo repository.BookRepository) BookService {
	return BookService{repo: repo}
}

func (s BookService) FindBooks() []domain.Book {
	return s.repo.GetAll()
}

func (s BookService) FindBook(isbn string) (domain.Book, error){
	return s.repo.GetByISBN(isbn)
}

func (s BookService) AddBook(b domain.Book) (domain.Book, error){
	return s.repo.Save(b)
}

func (s BookService) UpdateBook(b domain.Book) (domain.Book, error){
	return s.repo.Update(b)
}