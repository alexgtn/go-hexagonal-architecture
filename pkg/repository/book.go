package repository

import "github.com/alexgtn/esi2021-lab5/pkg/domain"

type BookRepository struct {
	Data []*domain.Book
}

func NewBookRepostory() *BookRepository {
	return &BookRepository{
		Data: []*domain.Book{},
	}
}

func (r *BookRepository) AddBook(s *domain.Book) (*domain.Book, error) {
	r.Data = append(r.Data, s)
	return s, nil
}

func (r *BookRepository) GetAllBooks() ([]*domain.Book, error) {
	return r.Data, nil
}
