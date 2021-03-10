package repository

import "github.com/alexgtn/esi2021-lab5/pkg/domain"

type StudentBarRepository struct {
	data []*domain.Student
}

func NewStudentBarRepostory() *StudentBarRepository {
	return &StudentBarRepository{
		data: []*domain.Student{},
	}
}

func (r *StudentBarRepository) AddStudent(s *domain.Student) (*domain.Student, error) {
	r.data = append(r.data, s)
	return s, nil
}

func (r *StudentBarRepository) GetAllStudents() ([]*domain.Student, error) {
	return r.data, nil
}
