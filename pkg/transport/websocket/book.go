package http

import (
	"encoding/json"
	"github.com/alexgtn/esi2021-lab5/pkg/domain"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type bookService interface {
	AddBook(bookmark *domain.Book) (*domain.Book, error)
	GetAllBooks() ([]*domain.Book, error)
}

type BookHandler struct {
	wsUpgrader  websocket.Upgrader
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
	c, err := h.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	book := &domain.Book{}
	err = json.Unmarshal(message, book)
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	createdBook, err := h.bookService.AddBook(book)
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	// write success response
	createdBookRes, err := json.Marshal(createdBook)
	err = c.WriteMessage(mt, createdBookRes)
	if err != nil {
		log.Println("write:", err)
	}
}

func (h *BookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	c, err := h.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(0, []byte(err.Error()))
		return
	}
	// write success response
	booksRes, err := json.Marshal(books)
	err = c.WriteMessage(0, booksRes)
	if err != nil {
		log.Println("write:", err)
	}
}
