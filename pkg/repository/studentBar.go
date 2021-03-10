package repository

import "github.com/alexgtn/esi2021-lab5/pkg/domain"

type StudentBarRepository struct {
	Data []*domain.Student
}

func NewStudentBarRepostory() *StudentBarRepository {
	return &StudentBarRepository{
		Data: []*domain.Student{},
	}
}

func (r *StudentBarRepository) AddStudent(s *domain.Student) (*domain.Student, error) {
	r.Data = append(r.Data, s)
	return s, nil
}

func (r *StudentBarRepository) GetAllStudents() ([]*domain.Student, error) {
	return r.Data, nil
}