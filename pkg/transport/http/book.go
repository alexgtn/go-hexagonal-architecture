package http

import (
"encoding/json"
"github.com/alexgtn/esi2021-lab5/pkg/domain"
"github.com/gorilla/mux"
log "github.com/sirupsen/logrus"
"net/http"
)

type bookService interface {
	AddBook(bookmark *domain.Book) (*domain.Book, error)
	GetAllBooks() ([]*domain.Book, error)
}

type BookHandler struct {
	bookService bookService
}

func NewBookHandler(bS bookService) *BookHandler {
	return &BookHandler{
		bookService: bS,
	}
}

func (h *BookHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/book", h.addBook).Methods(http.MethodPost)
	router.HandleFunc("/book", h.getAllBooks).Methods(http.MethodGet)
}

func (h *BookHandler) addBook(w http.ResponseWriter, r *http.Request) {

	book := &domain.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// close body to avoid memory leak
	err = r.Body.Close()
	if err != nil {
		log.Errorf("Could not close request body, err %v", err)
	}

	createdBook, err := h.bookService.AddBook(book)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&createdBook)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *BookHandler) getAllBooks(w http.ResponseWriter, _ *http.Request) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&books)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}
