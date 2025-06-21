package service

import (
	"course_schedule/internal/model"
	"time"
)

type ScheduleService interface {
	GetMySchedule(userID uint, userType string) ([]model.Section, error)
	GetTodaySchedule(userID uint, userType string) ([]model.Section, error)
	GetWeekSchedule(userID uint, userType string) ([]model.Section, error)
	GetStudentSchedule(studentID uint) ([]model.Section, error)
	GetTeacherSchedule(teacherID uint) ([]model.Section, error)
	ExportSchedule(userID uint, userType string) ([]model.Section, error)
}

type scheduleService struct {
	sectionService SectionService
}

func NewScheduleService(sectionService SectionService) ScheduleService {
	return &scheduleService{sectionService: sectionService}
}

func (s *scheduleService) GetMySchedule(userID uint, userType string) ([]model.Section, error) {
	switch userType {
	case "student":
		return s.sectionService.GetByStudent(userID)
	case "teacher":
		return s.sectionService.GetByTeacher(userID)
	default:
		return nil, nil
	}
}

func (s *scheduleService) GetTodaySchedule(userID uint, userType string) ([]model.Section, error) {
	sections, err := s.GetMySchedule(userID, userType)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	var todaySections []model.Section
	for _, sec := range sections {
		if sec.TimeSlot.DayOfWeek == weekday-1 {
			todaySections = append(todaySections, sec)
		}
	}
	return todaySections, nil
}

func (s *scheduleService) GetWeekSchedule(userID uint, userType string) ([]model.Section, error) {
	sections, err := s.GetMySchedule(userID, userType)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	year, week := now.ISOWeek()
	var weekSections []model.Section
	for _, sec := range sections {
		// 假设Section有StartDate字段，实际可根据业务调整
		if secStart, err := parseSectionStartDate(sec); err == nil {
			y, w := secStart.ISOWeek()
			if y == year && w == week {
				weekSections = append(weekSections, sec)
			}
		}
	}
	// 若无日期字段，暂时返回全部
	if len(weekSections) == 0 {
		return sections, nil
	}
	return weekSections, nil
}

func (s *scheduleService) GetStudentSchedule(studentID uint) ([]model.Section, error) {
	return s.sectionService.GetByStudent(studentID)
}

func (s *scheduleService) GetTeacherSchedule(teacherID uint) ([]model.Section, error) {
	return s.sectionService.GetByTeacher(teacherID)
}

func (s *scheduleService) ExportSchedule(userID uint, userType string) ([]model.Section, error) {
	// 目前直接复用GetMySchedule，后续可扩展格式/筛选
	return s.GetMySchedule(userID, userType)
}

// parseSectionStartDate 可根据实际Section结构实现
func parseSectionStartDate(sec model.Section) (time.Time, error) {
	// 示例：假设Section有StartDate string字段
	// return time.Parse("2006-01-02", sec.StartDate)
	return time.Time{}, nil // 占位
}
