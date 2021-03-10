package repository

import "github.com/alexgtn/esi2021-lab5/pkg/domain"

type StudentFooRepository struct {
	Data []*domain.Student
}

func NewStudentFooRepostory() *StudentFooRepository {
	return &StudentFooRepository{
		Data: []*domain.Student{},
	}
}

func (r *StudentFooRepository) AddStudent(s *domain.Student) (*domain.Student, error) {
	r.Data = append(r.Data, s)
	return s, nil
}

func (r *StudentFooRepository) GetAllStudents() ([]*domain.Student, error) {
	return r.Data, nil
}