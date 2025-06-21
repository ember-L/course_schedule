package handler

import (
	"course_schedule/internal/model"
	"course_schedule/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler 认证处理器
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler 创建认证处理器
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register 注册路由
func (h *AuthHandler) Register(e *echo.Echo) {
	// 公开路由 - 无需认证
	e.POST("/admin/login", h.AdminLogin)
	e.POST("/teacher/login", h.TeacherLogin)
	e.POST("/student/login", h.StudentLogin)
	e.POST("/logout", h.Logout)

	// 需要认证的路由
	authGroup := e.Group("")
	authGroup.Use(AuthMiddleware(h.authService))

	// 管理员路由
	adminGroup := authGroup.Group("/admin")
	adminGroup.Use(RoleMiddleware("admin"))
	adminGroup.GET("/profile/:id", h.GetAdminProfile)

	// 教师路由
	teacherGroup := authGroup.Group("/teacher")
	teacherGroup.Use(RoleMiddleware("teacher"))
	teacherGroup.GET("/profile/:id", h.GetTeacherProfile)

	// 学生路由
	studentGroup := authGroup.Group("/student")
	studentGroup.Use(RoleMiddleware("student"))
	studentGroup.GET("/profile/:id", h.GetStudentProfile)

	// 通用路由 - 所有认证用户都可以访问
	authGroup.POST("/change-password", h.ChangePassword)
}

// AdminLogin 管理员登录
func (h *AuthHandler) AdminLogin(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求参数"})
	}

	// 验证请求参数
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "用户名和密码不能为空"})
	}

	// 登录
	resp, err := h.authService.AdminLogin(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// GetAdminProfile 获取管理员信息
func (h *AuthHandler) GetAdminProfile(c echo.Context) error {
	// 从上下文中获取用户ID
	userID := getUserIDFromContext(c)

	// 获取管理员信息
	admin, err := h.authService.GetAdminProfile(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, admin)
}

// TeacherLogin 教师登录
func (h *AuthHandler) TeacherLogin(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求参数"})
	}

	// 验证请求参数
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "用户名和密码不能为空"})
	}

	// 登录
	resp, err := h.authService.TeacherLogin(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// GetTeacherProfile 获取教师信息
func (h *AuthHandler) GetTeacherProfile(c echo.Context) error {
	// 从上下文中获取用户ID
	userID := getUserIDFromContext(c)

	// 获取教师信息
	teacher, err := h.authService.GetTeacherProfile(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, teacher)
}

// StudentLogin 学生登录
func (h *AuthHandler) StudentLogin(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求参数"})
	}

	// 验证请求参数
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "用户名和密码不能为空"})
	}

	// 登录
	resp, err := h.authService.StudentLogin(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// GetStudentProfile 获取学生信息
func (h *AuthHandler) GetStudentProfile(c echo.Context) error {
	// 从上下文中获取用户ID
	userID := getUserIDFromContext(c)

	// 获取学生信息
	student, err := h.authService.GetStudentProfile(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}

// Logout 登出
func (h *AuthHandler) Logout(c echo.Context) error {
	// 客户端应该清除令牌，服务端无需处理
	return c.JSON(http.StatusOK, map[string]string{"message": "登出成功"})
}

// ChangePassword 修改密码
func (h *AuthHandler) ChangePassword(c echo.Context) error {
	var req model.ChangePasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求参数"})
	}

	// 验证请求参数
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "密码不能为空且新密码长度不能小于6位"})
	}

	// 从上下文中获取用户ID和类型
	userID := getUserIDFromContext(c)
	userType := getUserTypeFromContext(c)

	var err error
	switch userType {
	case "admin":
		err = h.authService.ChangeAdminPassword(userID, req.OldPassword, req.NewPassword)
	case "teacher":
		err = h.authService.ChangeTeacherPassword(userID, req.OldPassword, req.NewPassword)
	case "student":
		err = h.authService.ChangeStudentPassword(userID, req.OldPassword, req.NewPassword)
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的用户类型"})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "密码修改成功"})
}
