package repository

import (
	"course_schedule/internal/model"
	"errors"

	"gorm.io/gorm"
)

// SectionRepository 课程安排仓库接口
type SectionRepository interface {
	GetAll() ([]model.Section, error)
	GetByID(id uint) (*model.Section, error)
	Create(section *model.Section) error
	Update(section *model.Section) error
	Delete(id uint) error
	CheckConflict(section *model.Section) (bool, string, error)
}

// sectionRepository 课程安排仓库实现
type sectionRepository struct {
	db *Database
}

// NewSectionRepository 创建课程安排仓库
func NewSectionRepository(db *Database) SectionRepository {
	return &sectionRepository{db: db}
}

// GetAll 获取所有课程安排
func (r *sectionRepository) GetAll() ([]model.Section, error) {
	var sections []model.Section
	result := r.db.DB.
		Preload("Course").
		Preload("Teacher").
		Preload("Classroom").
		Preload("TimeSlot").
		Find(&sections)
	return sections, result.Error
}

// GetByID 通过ID获取课程安排
func (r *sectionRepository) GetByID(id uint) (*model.Section, error) {
	var section model.Section
	result := r.db.DB.
		Preload("Course").
		Preload("Teacher").
		Preload("Classroom").
		Preload("TimeSlot").
		First(&section, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &section, nil
}

// Create 创建课程安排
func (r *sectionRepository) Create(section *model.Section) error {
	// 先检查冲突
	hasConflict, _, err := r.CheckConflict(section)
	if err != nil {
		return err
	}
	if hasConflict {
		return errors.New("课程安排冲突")
	}
	return r.db.DB.Create(section).Error
}

// Update 更新课程安排
func (r *sectionRepository) Update(section *model.Section) error {
	// 先检查冲突
	hasConflict, _, err := r.CheckConflict(section)
	if err != nil {
		return err
	}
	if hasConflict {
		return errors.New("课程安排冲突")
	}
	return r.db.DB.Save(section).Error
}

// Delete 删除课程安排
func (r *sectionRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.Section{}, id).Error
}

// CheckConflict 检查课程安排冲突
func (r *sectionRepository) CheckConflict(section *model.Section) (bool, string, error) {
	// 检查教师时间冲突
	var teacherConflict model.Section
	teacherQuery := r.db.DB.Where("teacher_id = ? AND time_slot_id = ? AND semester = ?",
		section.TeacherID, section.TimeSlotID, section.Semester)
	
	// 如果是更新操作，排除自身
	if section.ID != 0 {
		teacherQuery = teacherQuery.Where("id != ?", section.ID)
	}
	
	result := teacherQuery.First(&teacherConflict)
	if result.Error == nil {
		// 找到冲突
		return true, "教师在该时间段已有其他课程安排", nil
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 查询出错
		return false, "", result.Error
	}
	
	// 检查教室时间冲突
	var classroomConflict model.Section
	classroomQuery := r.db.DB.Where("classroom_id = ? AND time_slot_id = ? AND semester = ?",
		section.ClassroomID, section.TimeSlotID, section.Semester)
	
	// 如果是更新操作，排除自身
	if section.ID != 0 {
		classroomQuery = classroomQuery.Where("id != ?", section.ID)
	}
	
	result = classroomQuery.First(&classroomConflict)
	if result.Error == nil {
		// 找到冲突
		return true, "教室在该时间段已被占用", nil
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 查询出错
		return false, "", result.Error
	}
	
	// 没有冲突
	return false, "", nil
}