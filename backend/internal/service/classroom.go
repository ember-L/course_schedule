package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

// ClassroomService 教室服务接口
type ClassroomService interface {
	GetAllClassrooms() ([]model.Classroom, error)
	GetClassroomByID(id uint) (*model.Classroom, error)
	CreateClassroom(classroom *model.Classroom) error
	UpdateClassroom(classroom *model.Classroom) error
	DeleteClassroom(id uint) error
}

// classroomService 教室服务实现
type classroomService struct {
	repo repository.ClassroomRepository
}

// NewClassroomService 创建教室服务
func NewClassroomService(repo repository.ClassroomRepository) ClassroomService {
	return &classroomService{repo: repo}
}

// GetAllClassrooms 获取所有教室
func (s *classroomService) GetAllClassrooms() ([]model.Classroom, error) {
	return s.repo.GetAll()
}

// GetClassroomByID 根据ID获取教室
func (s *classroomService) GetClassroomByID(id uint) (*model.Classroom, error) {
	return s.repo.GetByID(id)
}

// CreateClassroom 创建教室
func (s *classroomService) CreateClassroom(classroom *model.Classroom) error {
	return s.repo.Create(classroom)
}

// UpdateClassroom 更新教室
func (s *classroomService) UpdateClassroom(classroom *model.Classroom) error {
	return s.repo.Update(classroom)
}

// DeleteClassroom 删除教室
func (s *classroomService) DeleteClassroom(id uint) error {
	return s.repo.Delete(id)
}