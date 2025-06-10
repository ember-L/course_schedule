package repository

import (
	"course_schedule/internal/model"
)

// TeacherRepository 教师仓库接口
type TeacherRepository interface {
	GetAll() ([]model.Teacher, error)
	GetByID(id uint) (*model.Teacher, error)
	Create(teacher *model.Teacher) error
	Update(teacher *model.Teacher) error
	Delete(id uint) error
}

// teacherRepository 教师仓库实现
type teacherRepository struct {
	db *Database
}

// NewTeacherRepository 创建教师仓库
func NewTeacherRepository(db *Database) TeacherRepository {
	return &teacherRepository{db: db}
}

// GetAll 获取所有教师
func (r *teacherRepository) GetAll() ([]model.Teacher, error) {
	var teachers []model.Teacher
	result := r.db.DB.Find(&teachers)
	return teachers, result.Error
}

// GetByID 根据ID获取教师
func (r *teacherRepository) GetByID(id uint) (*model.Teacher, error) {
	var teacher model.Teacher
	result := r.db.DB.First(&teacher, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &teacher, nil
}

// Create 创建教师
func (r *teacherRepository) Create(teacher *model.Teacher) error {
	return r.db.DB.Create(teacher).Error
}

// Update 更新教师
func (r *teacherRepository) Update(teacher *model.Teacher) error {
	return r.db.DB.Save(teacher).Error
}

// Delete 删除教师
func (r *teacherRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.Teacher{}, id).Error
}