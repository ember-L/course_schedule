package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
	"fmt"
)

// SectionService 课程安排服务接口
type SectionService interface {
	GetAll() ([]model.Section, error)
	GetByID(id uint) (*model.Section, error)
	GetByTeacher(teacherId uint) ([]model.Section, error)
	GetByStudent(studentId uint) ([]model.Section, error)
	Create(section *model.Section) error
	Update(section *model.Section) error
	Delete(id uint) error
	CheckConflict(section *model.Section) (bool, string, error)
	AutoSchedule(req *model.AutoScheduleRequest, classroomService ClassroomService, timeSlotService TimeSlotService) (*model.Section, error)
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

// GetByTeacher 获取指定教师的课程安排
func (s *sectionService) GetByTeacher(teacherId uint) ([]model.Section, error) {
	return s.sectionRepo.GetByTeacher(teacherId)
}

// GetByStudent 获取指定学生的课程安排
func (s *sectionService) GetByStudent(studentId uint) ([]model.Section, error) {
	return s.sectionRepo.GetByStudent(studentId)
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

// AutoSchedule 自动排课实现
func (s *sectionService) AutoSchedule(req *model.AutoScheduleRequest, classroomService ClassroomService, timeSlotService TimeSlotService) (*model.Section, error) {
	// 获取所有教室和时间段
	classrooms, err := classroomService.GetAllClassrooms()
	if err != nil {
		return nil, err
	}
	timeSlots, err := timeSlotService.GetAllTimeSlots()
	if err != nil {
		return nil, err
	}

	// 遍历所有时间段和教室，查找无冲突的组合
	for _, slot := range timeSlots {
		for _, room := range classrooms {
			// 教室容量需满足要求
			if room.Capacity < req.MaxStudents {
				continue
			}
			section := &model.Section{
				CourseID:    req.CourseID,
				TeacherID:   req.TeacherID,
				ClassroomID: room.ID,
				TimeSlotID:  slot.ID,
				Semester:    req.Semester,
				MaxStudents: req.MaxStudents,
			}
			hasConflict, _, err := s.sectionRepo.CheckConflict(section)
			if err != nil {
				return nil, err
			}
			if !hasConflict {
				// 创建排课
				err := s.sectionRepo.Create(section)
				if err != nil {
					return nil, err
				}
				return section, nil
			}
		}
	}
	return nil, fmt.Errorf("无可用教室或时间段，无法自动排课")
}
