package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

// TimeSlotService 时间段服务接口
type TimeSlotService interface {
	GetAllTimeSlots() ([]model.TimeSlot, error)
	GetTimeSlotByID(id uint) (*model.TimeSlot, error)
	CreateTimeSlot(timeSlot *model.TimeSlot) error
	UpdateTimeSlot(timeSlot *model.TimeSlot) error
	DeleteTimeSlot(id uint) error
}

// timeSlotService 时间段服务实现
type timeSlotService struct {
	repo repository.TimeSlotRepository
}

// NewTimeSlotService 创建时间段服务
func NewTimeSlotService(repo repository.TimeSlotRepository) TimeSlotService {
	return &timeSlotService{repo: repo}
}

// GetAllTimeSlots 获取所有时间段
func (s *timeSlotService) GetAllTimeSlots() ([]model.TimeSlot, error) {
	return s.repo.GetAll()
}

// GetTimeSlotByID 根据ID获取时间段
func (s *timeSlotService) GetTimeSlotByID(id uint) (*model.TimeSlot, error) {
	return s.repo.GetByID(id)
}

// CreateTimeSlot 创建时间段
func (s *timeSlotService) CreateTimeSlot(timeSlot *model.TimeSlot) error {
	return s.repo.Create(timeSlot)
}

// UpdateTimeSlot 更新时间段
func (s *timeSlotService) UpdateTimeSlot(timeSlot *model.TimeSlot) error {
	return s.repo.Update(timeSlot)
}

// DeleteTimeSlot 删除时间段
func (s *timeSlotService) DeleteTimeSlot(id uint) error {
	return s.repo.Delete(id)
}