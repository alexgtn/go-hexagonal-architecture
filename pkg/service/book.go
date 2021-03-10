package service

import (
	"github.com/alexgtn/esi2021-lab5/pkg/domain"
)

type bookRepository interface {
	AddBook(s *domain.Book) (*domain.Book, error)
	GetAllBooks() ([]*domain.Book, error)
}

type BookService struct {
	bookRepository bookRepository
}

func NewBookService(bR bookRepository) *BookService {
	return &BookService{
		bookRepository: bR,
	}
}

func (s *BookService) AddBook(book *domain.Book) (*domain.Book, error) {
	return s.bookRepository.AddBook(book)
}

func (s *BookService) GetAllBooks() ([]*domain.Book, error){
	books, err := s.bookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}
	for _, book := range books {
		book.Title = translateToEstonian(book.Title)
	}
	return books, nil
}

func translateToEstonian(text string) string {
	// TODO
	return text
}