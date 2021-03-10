package repository

import "github.com/alexgtn/esi2021-lab5/pkg/domain"

type BookRepository struct {
	data []*domain.Book
}

func NewBookRepostory() *BookRepository {
	return &BookRepository{
		data: []*domain.Book{},
	}
}

func (r *BookRepository) AddBook(s *domain.Book) (*domain.Book, error) {
	r.data = append(r.data, s)
	return s, nil
}

func (r *BookRepository) GetAllBooks() ([]*domain.Book, error) {
	return r.data, nil
}
