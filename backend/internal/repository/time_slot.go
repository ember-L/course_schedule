package repository

import (
	"course_schedule/internal/model"
)

// TimeSlotRepository 时间段仓库接口
type TimeSlotRepository interface {
	GetAll() ([]model.TimeSlot, error)
	GetByID(id uint) (*model.TimeSlot, error)
	Create(timeSlot *model.TimeSlot) error
	Update(timeSlot *model.TimeSlot) error
	Delete(id uint) error
}

// timeSlotRepository 时间段仓库实现
type timeSlotRepository struct {
	db *Database
}

// NewTimeSlotRepository 创建时间段仓库
func NewTimeSlotRepository(db *Database) TimeSlotRepository {
	return &timeSlotRepository{db: db}
}

// GetAll 获取所有时间段
func (r *timeSlotRepository) GetAll() ([]model.TimeSlot, error) {
	var timeSlots []model.TimeSlot
	result := r.db.DB.Find(&timeSlots)
	return timeSlots, result.Error
}

// GetByID 根据ID获取时间段
func (r *timeSlotRepository) GetByID(id uint) (*model.TimeSlot, error) {
	var timeSlot model.TimeSlot
	result := r.db.DB.First(&timeSlot, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &timeSlot, nil
}

// Create 创建时间段
func (r *timeSlotRepository) Create(timeSlot *model.TimeSlot) error {
	return r.db.DB.Create(timeSlot).Error
}

// Update 更新时间段
func (r *timeSlotRepository) Update(timeSlot *model.TimeSlot) error {
	return r.db.DB.Save(timeSlot).Error
}

// Delete 删除时间段
func (r *timeSlotRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.TimeSlot{}, id).Error
}