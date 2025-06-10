package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

// TeacherService 教师服务接口
type TeacherService interface {
	GetAllTeachers() ([]model.Teacher, error)
	GetTeacherByID(id uint) (*model.Teacher, error)
	CreateTeacher(teacher *model.Teacher) error
	UpdateTeacher(teacher *model.Teacher) error
	DeleteTeacher(id uint) error
}

// teacherService 教师服务实现
type teacherService struct {
	repo repository.TeacherRepository
}

// NewTeacherService 创建教师服务
func NewTeacherService(repo repository.TeacherRepository) TeacherService {
	return &teacherService{repo: repo}
}

// GetAllTeachers 获取所有教师
func (s *teacherService) GetAllTeachers() ([]model.Teacher, error) {
	return s.repo.GetAll()
}

// GetTeacherByID 根据ID获取教师
func (s *teacherService) GetTeacherByID(id uint) (*model.Teacher, error) {
	return s.repo.GetByID(id)
}

// CreateTeacher 创建教师
func (s *teacherService) CreateTeacher(teacher *model.Teacher) error {
	return s.repo.Create(teacher)
}

// UpdateTeacher 更新教师
func (s *teacherService) UpdateTeacher(teacher *model.Teacher) error {
	return s.repo.Update(teacher)
}

// DeleteTeacher 删除教师
func (s *teacherService) DeleteTeacher(id uint) error {
	return s.repo.Delete(id)
}