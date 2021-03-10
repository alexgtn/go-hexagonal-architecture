package service

import (
	"github.com/alexgtn/esi2021-lab5/pkg/domain"
)

type studentRepository interface {
	AddStudent(s *domain.Student) (*domain.Student, error)
	GetAllStudents() ([]*domain.Student, error)
}

type StudentService struct {
	studentRepositoryFoo studentRepository
	studentRepositoryBar studentRepository
}

func NewStudentService(sRepoFoo studentRepository, sRepoBar studentRepository) *StudentService {
	return &StudentService{
		studentRepositoryFoo: sRepoFoo,
		studentRepositoryBar: sRepoBar,
	}
}

func (s *StudentService) AddStudent(student *domain.Student) (*domain.Student, error) {
	_, err := s.studentRepositoryFoo.AddStudent(student)
	if err != nil {
		return nil, err
	}
	_, err = s.studentRepositoryBar.AddStudent(student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) GetAllStudents() ([]*domain.Student, error){
	allStudents := []*domain.Student{}
	fooStudents, err := s.studentRepositoryFoo.GetAllStudents()
	if err != nil {
		return nil, err
	}
	barStudents, err := s.studentRepositoryBar.GetAllStudents()
	if err != nil {
		return nil, err
	}
	allStudents = append(allStudents, fooStudents...)
	allStudents = append(allStudents, barStudents...)
	return allStudents, nil
}