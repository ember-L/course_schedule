package repository

import (
	"course_schedule/internal/model"
)

// CourseRepository 课程仓库接口
type CourseRepository interface {
	GetAll() ([]model.Course, error)
	GetByID(id uint) (*model.Course, error)
	Create(course *model.Course) error
	Update(course *model.Course) error
	Delete(id uint) error
}

// courseRepository 课程仓库实现
type courseRepository struct {
	db *Database
}

// NewCourseRepository 创建课程仓库
func NewCourseRepository(db *Database) CourseRepository {
	return &courseRepository{db: db}
}

// GetAll 获取所有课程
func (r *courseRepository) GetAll() ([]model.Course, error) {
	var courses []model.Course
	result := r.db.DB.Find(&courses)
	return courses, result.Error
}

// GetByID 根据ID获取课程
func (r *courseRepository) GetByID(id uint) (*model.Course, error) {
	var course model.Course
	result := r.db.DB.First(&course, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}

// Create 创建课程
func (r *courseRepository) Create(course *model.Course) error {
	return r.db.DB.Create(course).Error
}

// Update 更新课程
func (r *courseRepository) Update(course *model.Course) error {
	return r.db.DB.Save(course).Error
}

// Delete 删除课程
func (r *courseRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.Course{}, id).Error
}