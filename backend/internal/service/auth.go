package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
	"course_schedule/pkg/auth"
	"errors"
)

// AuthService 认证服务接口
type AuthService interface {
	// 管理员认证
	AdminLogin(username, password string) (*model.LoginResponse, error)
	GetAdminProfile(id uint) (*model.AdminAuth, error)

	// 教师认证
	TeacherLogin(email, password string) (*model.LoginResponse, error)
	GetTeacherProfile(id uint) (*model.Teacher, error)

	// 学生认证
	StudentLogin(studentID, password string) (*model.LoginResponse, error)
	GetStudentProfile(id uint) (*model.Student, error)

	// 修改密码
	ChangeAdminPassword(id uint, oldPassword, newPassword string) error
	ChangeTeacherPassword(id uint, oldPassword, newPassword string) error
	ChangeStudentPassword(id uint, oldPassword, newPassword string) error

	// 验证令牌
	ValidateToken(token string) (*auth.JWTClaims, error)
}

// authService 认证服务实现
type authService struct {
	authRepo    repository.AuthRepository
	jwtSecret   string
	jwtDuration int
}

// NewAuthService 创建认证服务
func NewAuthService(authRepo repository.AuthRepository, jwtSecret string, jwtDuration int) AuthService {
	return &authService{
		authRepo:    authRepo,
		jwtSecret:   jwtSecret,
		jwtDuration: jwtDuration,
	}
}

// AdminLogin 管理员登录
func (s *authService) AdminLogin(username, password string) (*model.LoginResponse, error) {
	admin, err := s.authRepo.FindAdminByUsername(username)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !repository.CheckPasswordHash(password, admin.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成令牌
	token, expiresAt, err := auth.GenerateToken(admin.ID, admin.Username, "admin", s.jwtSecret, s.jwtDuration)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:     token,
		User:      admin,
		ExpiresAt: expiresAt,
	}, nil
}

// GetAdminProfile 获取管理员信息
func (s *authService) GetAdminProfile(id uint) (*model.AdminAuth, error) {
	admin, err := s.authRepo.FindAdminByUsername("")
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("管理员不存在")
	}
	return admin, nil
}

// TeacherLogin 教师登录
func (s *authService) TeacherLogin(email, password string) (*model.LoginResponse, error) {
	teacher, err := s.authRepo.FindTeacherByEmail(email)
	if err != nil {
		return nil, err
	}
	if teacher == nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !repository.CheckPasswordHash(password, teacher.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成令牌
	token, expiresAt, err := auth.GenerateToken(teacher.ID, email, "teacher", s.jwtSecret, s.jwtDuration)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:     token,
		User:      teacher,
		ExpiresAt: expiresAt,
	}, nil
}

// GetTeacherProfile 获取教师信息
func (s *authService) GetTeacherProfile(id uint) (*model.Teacher, error) {
	teacher, err := s.authRepo.FindTeacherByEmail("")
	if err != nil {
		return nil, err
	}
	if teacher == nil {
		return nil, errors.New("教师不存在")
	}
	return teacher, nil
}

// StudentLogin 学生登录
func (s *authService) StudentLogin(studentID, password string) (*model.LoginResponse, error) {
	student, err := s.authRepo.FindStudentByStudentID(studentID)
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !repository.CheckPasswordHash(password, student.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成令牌
	token, expiresAt, err := auth.GenerateToken(student.ID, studentID, "student", s.jwtSecret, s.jwtDuration)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:     token,
		User:      student,
		ExpiresAt: expiresAt,
	}, nil
}

// GetStudentProfile 获取学生信息
func (s *authService) GetStudentProfile(id uint) (*model.Student, error) {
	student, err := s.authRepo.FindStudentByStudentID("")
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, errors.New("学生不存在")
	}
	return student, nil
}

// ChangeAdminPassword 修改管理员密码
func (s *authService) ChangeAdminPassword(id uint, oldPassword, newPassword string) error {
	admin, err := s.authRepo.FindAdminByUsername("")
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New("管理员不存在")
	}

	// 验证旧密码
	if !repository.CheckPasswordHash(oldPassword, admin.Password) {
		return errors.New("原密码错误")
	}

	// 哈希新密码
	hashedPassword, err := repository.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新密码
	return s.authRepo.UpdateAdminPassword(id, hashedPassword)
}

// ChangeTeacherPassword 修改教师密码
func (s *authService) ChangeTeacherPassword(id uint, oldPassword, newPassword string) error {
	teacher, err := s.authRepo.FindTeacherByEmail("")
	if err != nil {
		return err
	}
	if teacher == nil {
		return errors.New("教师不存在")
	}

	// 验证旧密码
	if !repository.CheckPasswordHash(oldPassword, teacher.Password) {
		return errors.New("原密码错误")
	}

	// 哈希新密码
	hashedPassword, err := repository.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新密码
	return s.authRepo.UpdateTeacherPassword(id, hashedPassword)
}

// ChangeStudentPassword 修改学生密码
func (s *authService) ChangeStudentPassword(id uint, oldPassword, newPassword string) error {
	student, err := s.authRepo.FindStudentByStudentID("")
	if err != nil {
		return err
	}
	if student == nil {
		return errors.New("学生不存在")
	}

	// 验证旧密码
	if !repository.CheckPasswordHash(oldPassword, student.Password) {
		return errors.New("原密码错误")
	}

	// 哈希新密码
	hashedPassword, err := repository.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新密码
	return s.authRepo.UpdateStudentPassword(id, hashedPassword)
}

// ValidateToken 验证令牌
func (s *authService) ValidateToken(token string) (*auth.JWTClaims, error) {
	return auth.ValidateToken(token, s.jwtSecret)
}