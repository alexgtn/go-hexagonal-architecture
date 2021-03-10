package repository

import "github.com/alexgtn/esi2021-lab5/pkg/domain"

type StudentFooRepository struct {
	data []*domain.Student
}

func NewStudentFooRepostory() *StudentFooRepository {
	return &StudentFooRepository{
		data: []*domain.Student{},
	}
}

func (r *StudentFooRepository) AddStudent(s *domain.Student) (*domain.Student, error) {
	r.data = append(r.data, s)
	return s, nil
}

func (r *StudentFooRepository) GetAllStudents() ([]*domain.Student, error) {
	return r.data, nil
}
