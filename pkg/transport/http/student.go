package http

import (
	"encoding/json"
	"github.com/alexgtn/esi2021-lab5/pkg/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type studentService interface {
	AddStudent(bookmark *domain.Student) (*domain.Student, error)
	GetAllStudents() ([]*domain.Student, error)
}

type StudentHandler struct {
	bookService studentService
}

func NewStudentHandler(bS studentService) *StudentHandler {
	return &StudentHandler{
		bookService: bS,
	}
}

func (h *StudentHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/student", h.addStudent).Methods(http.MethodPost)
	router.HandleFunc("/student", h.getAllStudents).Methods(http.MethodGet)
}

func (h *StudentHandler) addStudent(w http.ResponseWriter, r *http.Request) {
	book := &domain.Student{}
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

	createdStudent, err := h.bookService.AddStudent(book)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&createdStudent)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *StudentHandler) getAllStudents(w http.ResponseWriter, _ *http.Request) {
	bookmarks, err := h.bookService.GetAllStudents()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&bookmarks)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}
