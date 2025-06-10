package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 基础模型
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TeacherCourse 教师课程关联模型
type TeacherCourse struct {
	TeacherID uint `gorm:"primaryKey"`
	CourseID  uint `gorm:"primaryKey"`
}

// Teacher 教师模型
type Teacher struct {
	BaseModel
	Name     string    `json:"name" gorm:"size:50;not null;index"`
	Title    string    `json:"title" gorm:"size:50"`
	Email    string    `json:"email" gorm:"size:100;uniqueIndex;not null"`
	Password string    `json:"-" gorm:"size:100;not null;default:'$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'"`
	Courses  []Course  `json:"courses,omitempty" gorm:"many2many:teacher_courses;joinForeignKey:teacher_id;joinReferences:course_id"`
	Sections []Section `json:"sections,omitempty"`
}

// Course 课程模型
type Course struct {
	BaseModel
	Code        string    `json:"code" gorm:"size:20;uniqueIndex;not null"`
	Name        string    `json:"name" gorm:"size:100;not null;index"`
	Credits     float32   `json:"credits" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	Teachers    []Teacher `json:"teachers,omitempty" gorm:"many2many:teacher_courses;joinForeignKey:course_id;joinReferences:teacher_id"`
	Sections    []Section `json:"sections,omitempty"`
}

// Classroom 教室模型
type Classroom struct {
	BaseModel
	Building string    `json:"building" gorm:"size:50;not null;index"`
	Room     string    `json:"room" gorm:"size:20;not null"`
	Capacity int       `json:"capacity" gorm:"not null"`
	Type     string    `json:"type" gorm:"size:50;index"` // 普通教室、实验室、多媒体教室等
	Sections []Section `json:"sections,omitempty"`
}

// TimeSlot 时间段模型
type TimeSlot struct {
	BaseModel
	DayOfWeek int       `json:"dayOfWeek" gorm:"not null;index"`  // 1-7 表示周一到周日
	StartTime string    `json:"startTime" gorm:"size:5;not null"` // 格式如 "08:00"
	EndTime   string    `json:"endTime" gorm:"size:5;not null"`   // 格式如 "09:40"
	Sections  []Section `json:"sections,omitempty"`
}

// Section 课程安排模型
type Section struct {
	BaseModel
	CourseID    uint         `json:"courseId" gorm:"not null;index"`
	Course      Course       `json:"course,omitempty"`
	TeacherID   uint         `json:"teacherId" gorm:"not null;index"`
	Teacher     Teacher      `json:"teacher,omitempty"`
	ClassroomID uint         `json:"classroomId" gorm:"not null;index"`
	Classroom   Classroom    `json:"classroom,omitempty"`
	TimeSlotID  uint         `json:"timeSlotId" gorm:"not null;index"`
	TimeSlot    TimeSlot     `json:"timeSlot,omitempty"`
	Semester    string       `json:"semester" gorm:"size:20;not null;index"` // 如 "2023-2024-1"
	MaxStudents int          `json:"maxStudents" gorm:"not null"`
	Enrollments []Enrollment `json:"enrollments,omitempty"`
}

// Student 学生模型
type Student struct {
	BaseModel
	StudentID   string       `json:"studentId" gorm:"size:20;uniqueIndex;not null"`
	Name        string       `json:"name" gorm:"size:50;not null;index"`
	Grade       string       `json:"grade" gorm:"size:10;not null;index"` // 年级
	Major       string       `json:"major" gorm:"size:50;not null;index"` // 专业
	Password    string       `json:"-" gorm:"size:100;not null;default:'$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'"`
	Enrollments []Enrollment `json:"enrollments,omitempty"`
}

// Enrollment 选课记录模型
type Enrollment struct {
	BaseModel
	StudentID  uint      `json:"studentId" gorm:"not null;index"`
	Student    Student   `json:"student,omitempty"`
	SectionID  uint      `json:"sectionId" gorm:"not null;index"`
	Section    Section   `json:"section,omitempty"`
	EnrollTime time.Time `json:"enrollTime" gorm:"not null"`
	Grade      string    `json:"grade" gorm:"size:2"` // 成绩
}

// AdminAuth 管理员认证信息
type AdminAuth struct {
	BaseModel
	Username string `json:"username" gorm:"size:50;uniqueIndex;not null"`
	Password string `json:"-" gorm:"size:100;not null"`
	Name     string `json:"name" gorm:"size:50;not null"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string      `json:"token"`
	User      interface{} `json:"user"`
	ExpiresAt time.Time   `json:"expiresAt"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}
