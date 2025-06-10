package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

// StudentService 学生服务接口
type StudentService interface {
	GetAllStudents() ([]model.Student, error)
	GetStudentByID(id uint) (*model.Student, error)
	GetStudentByStudentID(studentID string) (*model.Student, error)
	CreateStudent(student *model.Student) error
	UpdateStudent(student *model.Student) error
	DeleteStudent(id uint) error
}

// studentService 学生服务实现
type studentService struct {
	repo repository.StudentRepository
}

// NewStudentService 创建学生服务
func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

// GetAllStudents 获取所有学生
func (s *studentService) GetAllStudents() ([]model.Student, error) {
	return s.repo.GetAll()
}

// GetStudentByID 根据ID获取学生
func (s *studentService) GetStudentByID(id uint) (*model.Student, error) {
	return s.repo.GetByID(id)
}

// GetStudentByStudentID 根据学号获取学生
func (s *studentService) GetStudentByStudentID(studentID string) (*model.Student, error) {
	return s.repo.GetByStudentID(studentID)
}

// CreateStudent 创建学生
func (s *studentService) CreateStudent(student *model.Student) error {
	return s.repo.Create(student)
}

// UpdateStudent 更新学生
func (s *studentService) UpdateStudent(student *model.Student) error {
	return s.repo.Update(student)
}

// DeleteStudent 删除学生
func (s *studentService) DeleteStudent(id uint) error {
	return s.repo.Delete(id)
}