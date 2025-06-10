package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

// CourseService 课程服务接口
type CourseService interface {
	GetAllCourses() ([]model.Course, error)
	GetCourseByID(id uint) (*model.Course, error)
	CreateCourse(course *model.Course) error
	UpdateCourse(course *model.Course) error
	DeleteCourse(id uint) error
}

// courseService 课程服务实现
type courseService struct {
	repo repository.CourseRepository
}

// NewCourseService 创建课程服务
func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

// GetAllCourses 获取所有课程
func (s *courseService) GetAllCourses() ([]model.Course, error) {
	return s.repo.GetAll()
}

// GetCourseByID 根据ID获取课程
func (s *courseService) GetCourseByID(id uint) (*model.Course, error) {
	return s.repo.GetByID(id)
}

// CreateCourse 创建课程
func (s *courseService) CreateCourse(course *model.Course) error {
	return s.repo.Create(course)
}

// UpdateCourse 更新课程
func (s *courseService) UpdateCourse(course *model.Course) error {
	return s.repo.Update(course)
}

// DeleteCourse 删除课程
func (s *courseService) DeleteCourse(id uint) error {
	return s.repo.Delete(id)
}