package http

import (
"encoding/json"
"github.com/alexgtn/esi2021-lab5/pkg/domain"
"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
"net/http"
)

type studentService interface {
	AddStudent(studentmark *domain.Student) (*domain.Student, error)
	GetAllStudents() ([]*domain.Student, error)
}

type StudentHandler struct {
	wsUpgrader websocket.Upgrader
	studentService studentService
}

func NewStudentHandler(bS studentService) *StudentHandler {
	return &StudentHandler{
		studentService: bS,
	}
}

func (h *StudentHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/student", h.addStudent).Methods(http.MethodPost)
	router.HandleFunc("/student", h.getAllStudents).Methods(http.MethodGet)
}

func (h *StudentHandler) addStudent(w http.ResponseWriter, r *http.Request) {
	c, err := h.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer c.Close()
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	student := &domain.Student{}
	err = json.Unmarshal(message, student)
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	createdStudent, err := h.studentService.AddStudent(student)
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	// write success response
	createdStudentRes, err := json.Marshal(createdStudent)
	err = c.WriteMessage(mt, createdStudentRes)
	if err != nil {
		log.Println("write:", err)
	}
}

func (h *StudentHandler) getAllStudents(w http.ResponseWriter, r *http.Request) {
	c, err := h.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	students, err := h.studentService.GetAllStudents()
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(0, []byte(err.Error()))
		return
	}
	// write success response
	studentsRes, err := json.Marshal(students)
	err = c.WriteMessage(0, studentsRes)
	if err != nil {
		log.Println("write:", err)
	}
}
