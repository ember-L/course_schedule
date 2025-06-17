package repository

import (
	"course_schedule/internal/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthRepository 认证仓库接口
type AuthRepository interface {
	// 管理员认证
	FindAdminByUsername(username string) (*model.AdminAuth, error)
	GetAdminById(id uint) (*model.AdminAuth, error)
	CreateAdmin(admin *model.AdminAuth) error
	UpdateAdminPassword(id uint, hashedPassword string) error

	// 教师认证
	FindTeacherByEmail(email string) (*model.Teacher, error)
	GetTeacherById(id uint) (*model.Teacher, error)
	UpdateTeacherPassword(id uint, hashedPassword string) error

	// 学生认证
	FindStudentByStudentID(studentID string) (*model.Student, error)
	GetStudentById(id uint) (*model.Student, error)
	UpdateStudentPassword(id uint, hashedPassword string) error
}

// authRepository 认证仓库实现
type authRepository struct {
	db *Database
}

// NewAuthRepository 创建认证仓库
func NewAuthRepository(db *Database) AuthRepository {
	return &authRepository{db: db}
}

// FindAdminByUsername 通过用户名查找管理员
func (r *authRepository) FindAdminByUsername(username string) (*model.AdminAuth, error) {
	var admin model.AdminAuth
	result := r.db.DB.Where("username = ?", username).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &admin, nil
}

func (r *authRepository) GetAdminById(id uint) (*model.AdminAuth, error) {
	var admin model.AdminAuth
	result := r.db.DB.Where("id = ?", id).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &admin, nil
}

// CreateAdmin 创建管理员
func (r *authRepository) CreateAdmin(admin *model.AdminAuth) error {
	return r.db.DB.Create(admin).Error
}

// UpdateAdminPassword 更新管理员密码
func (r *authRepository) UpdateAdminPassword(id uint, hashedPassword string) error {
	return r.db.DB.Model(&model.AdminAuth{}).Where("id = ?", id).Update("password", hashedPassword).Error
}

// FindTeacherByEmail 通过邮箱查找教师
func (r *authRepository) FindTeacherByEmail(email string) (*model.Teacher, error) {
	var teacher model.Teacher
	result := r.db.DB.Where("email = ?", email).First(&teacher)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &teacher, nil
}

func (r *authRepository) GetTeacherById(id uint) (*model.Teacher, error) {
	var teacher model.Teacher
	result := r.db.DB.Where("id = ?", id).First(&teacher)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &teacher, nil
}

// UpdateTeacherPassword 更新教师密码
func (r *authRepository) UpdateTeacherPassword(id uint, hashedPassword string) error {
	return r.db.DB.Model(&model.Teacher{}).Where("id = ?", id).Update("password", hashedPassword).Error
}

// FindStudentByStudentID 通过学号查找学生
func (r *authRepository) FindStudentByStudentID(studentID string) (*model.Student, error) {
	var student model.Student
	result := r.db.DB.Where("student_id = ?", studentID).First(&student)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &student, nil
}

func (r *authRepository) GetStudentById(id uint) (*model.Student, error) {
	var student model.Student
	result := r.db.DB.Where("id = ?", id).First(&student)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &student, nil
}

// UpdateStudentPassword 更新学生密码
func (r *authRepository) UpdateStudentPassword(id uint, hashedPassword string) error {
	return r.db.DB.Model(&model.Student{}).Where("id = ?", id).Update("password", hashedPassword).Error
}

// HashPassword 密码哈希
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 检查密码哈希
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
