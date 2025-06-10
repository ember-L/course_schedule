package service

import (
	"errors"
	"time"
	
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

// EnrollmentService 选课服务接口
type EnrollmentService interface {
	GetAllEnrollments() ([]model.Enrollment, error)
	GetEnrollmentByID(id uint) (*model.Enrollment, error)
	GetEnrollmentsByStudentID(studentID uint) ([]model.Enrollment, error)
	GetEnrollmentsBySectionID(sectionID uint) ([]model.Enrollment, error)
	EnrollCourse(studentID, sectionID uint) error
	UpdateEnrollment(enrollment *model.Enrollment) error
	DeleteEnrollment(id uint) error
}

// enrollmentService 选课服务实现
type enrollmentService struct {
	enrollRepo repository.EnrollmentRepository
	sectionRepo repository.SectionRepository
}

// NewEnrollmentService 创建选课服务
func NewEnrollmentService(enrollRepo repository.EnrollmentRepository, sectionRepo repository.SectionRepository) EnrollmentService {
	return &enrollmentService{
		enrollRepo: enrollRepo,
		sectionRepo: sectionRepo,
	}
}

// GetAllEnrollments 获取所有选课记录
func (s *enrollmentService) GetAllEnrollments() ([]model.Enrollment, error) {
	return s.enrollRepo.GetAll()
}

// GetEnrollmentByID 根据ID获取选课记录
func (s *enrollmentService) GetEnrollmentByID(id uint) (*model.Enrollment, error) {
	return s.enrollRepo.GetByID(id)
}

// GetEnrollmentsByStudentID 根据学生ID获取选课记录
func (s *enrollmentService) GetEnrollmentsByStudentID(studentID uint) ([]model.Enrollment, error) {
	return s.enrollRepo.GetByStudentID(studentID)
}

// GetEnrollmentsBySectionID 根据课程安排ID获取选课记录
func (s *enrollmentService) GetEnrollmentsBySectionID(sectionID uint) ([]model.Enrollment, error) {
	return s.enrollRepo.GetBySectionID(sectionID)
}

// EnrollCourse 学生选课
func (s *enrollmentService) EnrollCourse(studentID, sectionID uint) error {
	// 检查是否已选过该课程
	exists, err := s.enrollRepo.CheckEnrollment(studentID, sectionID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("已经选过该课程")
	}
	
	// 检查课程是否已满
	section, err := s.sectionRepo.GetByID(sectionID)
	if err != nil {
		return err
	}
	
	enrollments, err := s.enrollRepo.GetBySectionID(sectionID)
	if err != nil {
		return err
	}
	
	if len(enrollments) >= section.MaxStudents {
		return errors.New("课程已满")
	}
	
	// 创建选课记录
	enrollment := &model.Enrollment{
		StudentID:  studentID,
		SectionID:  sectionID,
		EnrollTime: time.Now(),
	}
	
	return s.enrollRepo.Create(enrollment)
}

// UpdateEnrollment 更新选课记录
func (s *enrollmentService) UpdateEnrollment(enrollment *model.Enrollment) error {
	return s.enrollRepo.Update(enrollment)
}

// DeleteEnrollment 删除选课记录
func (s *enrollmentService) DeleteEnrollment(id uint) error {
	return s.enrollRepo.Delete(id)
}