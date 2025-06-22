package repository

import (
	"course_schedule/internal/model"
	"fmt"
)

// EnrollmentRepository 选课记录仓库接口
type EnrollmentRepository interface {
	GetAll() ([]model.Enrollment, error)
	GetByID(id uint) (*model.Enrollment, error)
	GetByStudentID(studentID uint) ([]model.Enrollment, error)
	GetBySectionID(sectionID uint) ([]model.Enrollment, error)
	GetByTeacherID(teacherID uint) ([]model.Enrollment, error)
	GetByTeacherAndSection(teacherID, sectionID uint) ([]model.Enrollment, error)
	IsTeacherCourse(teacherID, sectionID uint) (bool, error)
	Create(enrollment *model.Enrollment) error
	Update(enrollment *model.Enrollment) error
	Delete(id uint) error
	CheckEnrollment(studentID, sectionID uint) (bool, error)
}

// enrollmentRepository 选课记录仓库实现
type enrollmentRepository struct {
	db *Database
}

// NewEnrollmentRepository 创建选课记录仓库
func NewEnrollmentRepository(db *Database) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

// GetAll 获取所有选课记录
func (r *enrollmentRepository) GetAll() ([]model.Enrollment, error) {
	var enrollments []model.Enrollment
	result := r.db.DB.Preload("Student").Preload("Section").Find(&enrollments)
	return enrollments, result.Error
}

// GetByID 根据ID获取选课记录
func (r *enrollmentRepository) GetByID(id uint) (*model.Enrollment, error) {
	var enrollment model.Enrollment
	result := r.db.DB.Preload("Student").Preload("Section").First(&enrollment, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &enrollment, nil
}

// GetByStudentID 根据学生ID获取选课记录
func (r *enrollmentRepository) GetByStudentID(studentID uint) ([]model.Enrollment, error) {
	var enrollments []model.Enrollment
	result := r.db.DB.Where("student_id = ?", studentID).
		Preload("Student").Preload("Section").
		Preload("Section.Course").Preload("Section.Teacher").
		Preload("Section.Classroom").Preload("Section.TimeSlot").
		Find(&enrollments)
	return enrollments, result.Error
}

// GetBySectionID 根据课程安排ID获取选课记录
func (r *enrollmentRepository) GetBySectionID(sectionID uint) ([]model.Enrollment, error) {
	var enrollments []model.Enrollment
	result := r.db.DB.Where("section_id = ?", sectionID).
		Preload("Student").Preload("Section").
		Find(&enrollments)
	return enrollments, result.Error
}

// GetByTeacherID 根据教师ID获取该教师所有课程的选课记录
func (r *enrollmentRepository) GetByTeacherID(teacherID uint) ([]model.Enrollment, error) {
	var enrollments []model.Enrollment
	result := r.db.DB.
		Joins("JOIN section ON enrollment.section_id = section.id").
		Where("section.teacher_id = ?", teacherID).
		Preload("Student").Preload("Section").
		Preload("Section.Course").Preload("Section.Teacher").
		Preload("Section.Classroom").Preload("Section.TimeSlot").
		Find(&enrollments)
	return enrollments, result.Error
}

// GetByTeacherAndSection 根据教师ID和课程安排ID获取选课记录
func (r *enrollmentRepository) GetByTeacherAndSection(teacherID, sectionID uint) ([]model.Enrollment, error) {
	var enrollments []model.Enrollment
	result := r.db.DB.
		Joins("JOIN section ON enrollment.section_id = section.id").
		Where("section.teacher_id = ? AND enrollment.section_id = ?", teacherID, sectionID).
		Preload("Student").Preload("Section").
		Preload("Section.Course").Preload("Section.Teacher").
		Preload("Section.Classroom").Preload("Section.TimeSlot").
		Find(&enrollments)
	return enrollments, result.Error
}

// IsTeacherCourse 检查课程是否属于指定教师
func (r *enrollmentRepository) IsTeacherCourse(teacherID, sectionID uint) (bool, error) {
	var count int64
	err := r.db.DB.Model(&model.Section{}).
		Where("id = ? AND teacher_id = ?", sectionID, teacherID).
		Count(&count).Error

	return count > 0, err
}

// Create 创建选课记录
func (r *enrollmentRepository) Create(enrollment *model.Enrollment) error {
	return r.db.DB.Create(enrollment).Error
}

// Update 更新选课记录
func (r *enrollmentRepository) Update(enrollment *model.Enrollment) error {
	fmt.Printf("Updating enrollment:%v", *enrollment)
	return r.db.DB.Model(&enrollment).Update("grade", enrollment.Grade).Error
}

// Delete 删除选课记录
func (r *enrollmentRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.Enrollment{}, id).Error
}

// CheckEnrollment 检查学生是否已选该课程
func (r *enrollmentRepository) CheckEnrollment(studentID, sectionID uint) (bool, error) {
	var count int64
	err := r.db.DB.Model(&model.Enrollment{}).
		Where("student_id = ? AND section_id = ?", studentID, sectionID).
		Count(&count).Error

	return count > 0, err
}
