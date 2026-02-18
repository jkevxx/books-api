package service

import (
	"books-api/internal/model"
	"books-api/internal/store"
	"errors"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) GetAllBooks() ([]*model.Book, error) {
	return s.store.GetAll()
}

func (s *Service) GetBookById(id int) (*model.Book, error) {
	return s.store.GetById(id)
}

func (s *Service) CreateBook(book *model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("title is required")
	}

	return s.store.Create(book)
}

func (s *Service) UpdateBook(id int, book *model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("title is required")
	}

	return s.store.Update(id, book)
}

func (s *Service) DeleteBook(id int) error {
	return s.store.Delete(id)
}
