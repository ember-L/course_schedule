package repository

import (
	"course_schedule/internal/model"
)

// StudentRepository 学生仓库接口
type StudentRepository interface {
	GetAll() ([]model.Student, error)
	GetByID(id uint) (*model.Student, error)
	GetByStudentID(studentID string) (*model.Student, error)
	Create(student *model.Student) error
	Update(student *model.Student) error
	Delete(id uint) error
}

// studentRepository 学生仓库实现
type studentRepository struct {
	db *Database
}

// NewStudentRepository 创建学生仓库
func NewStudentRepository(db *Database) StudentRepository {
	return &studentRepository{db: db}
}

// GetAll 获取所有学生
func (r *studentRepository) GetAll() ([]model.Student, error) {
	var students []model.Student
	result := r.db.DB.Find(&students)
	return students, result.Error
}

// GetByID 根据ID获取学生
func (r *studentRepository) GetByID(id uint) (*model.Student, error) {
	var student model.Student
	result := r.db.DB.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

// GetByStudentID 根据学号获取学生
func (r *studentRepository) GetByStudentID(studentID string) (*model.Student, error) {
	var student model.Student
	result := r.db.DB.Where("student_id = ?", studentID).First(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

// Create 创建学生
func (r *studentRepository) Create(student *model.Student) error {
	return r.db.DB.Create(student).Error
}

// Update 更新学生
func (r *studentRepository) Update(student *model.Student) error {
	return r.db.DB.Save(student).Error
}

// Delete 删除学生
func (r *studentRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.Student{}, id).Error
}