package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

// SectionService 课程安排服务接口
type SectionService interface {
	GetAll() ([]model.Section, error)
	GetByID(id uint) (*model.Section, error)
	Create(section *model.Section) error
	Update(section *model.Section) error
	Delete(id uint) error
	CheckConflict(section *model.Section) (bool, string, error)
}

// sectionService 课程安排服务实现
type sectionService struct {
	sectionRepo repository.SectionRepository
}

// NewSectionService 创建课程安排服务
func NewSectionService(sectionRepo repository.SectionRepository) SectionService {
	return &sectionService{
		sectionRepo: sectionRepo,
	}
}

// GetAll 获取所有课程安排
func (s *sectionService) GetAll() ([]model.Section, error) {
	return s.sectionRepo.GetAll()
}

// GetByID 通过ID获取课程安排
func (s *sectionService) GetByID(id uint) (*model.Section, error) {
	return s.sectionRepo.GetByID(id)
}

// Create 创建课程安排
func (s *sectionService) Create(section *model.Section) error {
	return s.sectionRepo.Create(section)
}

// Update 更新课程安排
func (s *sectionService) Update(section *model.Section) error {
	return s.sectionRepo.Update(section)
}

// Delete 删除课程安排
func (s *sectionService) Delete(id uint) error {
	return s.sectionRepo.Delete(id)
}

// CheckConflict 检查课程安排冲突
func (s *sectionService) CheckConflict(section *model.Section) (bool, string, error) {
	return s.sectionRepo.CheckConflict(section)
}