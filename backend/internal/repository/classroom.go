package repository

import (
	"course_schedule/internal/model"
)

// ClassroomRepository 教室仓库接口
type ClassroomRepository interface {
	GetAll() ([]model.Classroom, error)
	GetByID(id uint) (*model.Classroom, error)
	Create(classroom *model.Classroom) error
	Update(classroom *model.Classroom) error
	Delete(id uint) error
}

// classroomRepository 教室仓库实现
type classroomRepository struct {
	db *Database
}

// NewClassroomRepository 创建教室仓库
func NewClassroomRepository(db *Database) ClassroomRepository {
	return &classroomRepository{db: db}
}

// GetAll 获取所有教室
func (r *classroomRepository) GetAll() ([]model.Classroom, error) {
	var classrooms []model.Classroom
	result := r.db.DB.Find(&classrooms)
	return classrooms, result.Error
}

// GetByID 根据ID获取教室
func (r *classroomRepository) GetByID(id uint) (*model.Classroom, error) {
	var classroom model.Classroom
	result := r.db.DB.First(&classroom, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &classroom, nil
}

// Create 创建教室
func (r *classroomRepository) Create(classroom *model.Classroom) error {
	return r.db.DB.Create(classroom).Error
}

// Update 更新教室
func (r *classroomRepository) Update(classroom *model.Classroom) error {
	return r.db.DB.Save(classroom).Error
}

// Delete 删除教室
func (r *classroomRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.Classroom{}, id).Error
}