package handler

import (
	"net/http"
	"strconv"

	"course_schedule/internal/model"
	"course_schedule/internal/service"

	"github.com/labstack/echo/v4"
)

// EnrollmentHandler 选课记录处理器
type EnrollmentHandler struct {
	enrollmentService service.EnrollmentService
	authService       service.AuthService
}

// NewEnrollmentHandler 创建选课记录处理器
func NewEnrollmentHandler(enrollmentService service.EnrollmentService, authService service.AuthService) *EnrollmentHandler {
	return &EnrollmentHandler{
		enrollmentService: enrollmentService,
		authService:       authService,
	}
}

// Register 注册路由
func (h *EnrollmentHandler) Register(e *echo.Echo) {
	// 需要认证的路由
	authGroup := e.Group("")
	authGroup.Use(AuthMiddleware(h.authService))

	// 管理员可以访问所有选课记录
	adminGroup := authGroup.Group("/enrollments")
	adminGroup.Use(RoleMiddleware("admin"))
	adminGroup.GET("", h.GetAllEnrollments)
	adminGroup.GET("/:id", h.GetEnrollmentByID)
	adminGroup.GET("/student/:studentId", h.GetEnrollmentsByStudentID)
	adminGroup.GET("/section/:sectionId", h.GetEnrollmentsBySectionID)
	adminGroup.POST("", h.EnrollCourse)
	adminGroup.PUT("/:id", h.UpdateEnrollment)
	adminGroup.DELETE("/:id", h.DeleteEnrollment)

	// 教师只能查看自己课程学生的选课记录
	teacherGroup := authGroup.Group("/enrollments")
	teacherGroup.Use(RoleMiddleware("teacher"))
	teacherGroup.GET("/my-courses", h.GetMyCourseEnrollments)
	teacherGroup.GET("/my-courses/:sectionId", h.GetMySectionEnrollments)
	teacherGroup.PUT("/:id/grade", h.UpdateEnrollmentGrade)

	// 学生只能查看自己的选课记录
	studentGroup := authGroup.Group("/enrollments")
	studentGroup.Use(RoleMiddleware("student"))
	studentGroup.GET("/my", h.GetMyEnrollments)
	studentGroup.POST("/enroll", h.EnrollCourse)
	studentGroup.DELETE("/drop/:id", h.DropCourse)
}

// GetAllEnrollments 获取所有选课记录（仅管理员）
func (h *EnrollmentHandler) GetAllEnrollments(c echo.Context) error {
	enrollments, err := h.enrollmentService.GetAllEnrollments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enrollments)
}

// GetEnrollmentByID 根据ID获取选课记录（仅管理员）
func (h *EnrollmentHandler) GetEnrollmentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	enrollment, err := h.enrollmentService.GetEnrollmentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Enrollment not found"})
	}

	return c.JSON(http.StatusOK, enrollment)
}

// GetEnrollmentsByStudentID 根据学生ID获取选课记录（仅管理员）
func (h *EnrollmentHandler) GetEnrollmentsByStudentID(c echo.Context) error {
	studentID, err := strconv.ParseUint(c.Param("studentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	enrollments, err := h.enrollmentService.GetEnrollmentsByStudentID(uint(studentID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enrollments)
}

// GetEnrollmentsBySectionID 根据课程安排ID获取选课记录（仅管理员）
func (h *EnrollmentHandler) GetEnrollmentsBySectionID(c echo.Context) error {
	sectionID, err := strconv.ParseUint(c.Param("sectionId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid section ID"})
	}

	enrollments, err := h.enrollmentService.GetEnrollmentsBySectionID(uint(sectionID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enrollments)
}

// GetMyCourseEnrollments 获取教师自己课程的所有选课记录
func (h *EnrollmentHandler) GetMyCourseEnrollments(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}

	enrollments, err := h.enrollmentService.GetEnrollmentsByTeacherID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enrollments)
}

// GetMySectionEnrollments 获取教师指定课程的学生选课记录
func (h *EnrollmentHandler) GetMySectionEnrollments(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}

	sectionID, err := strconv.ParseUint(c.Param("sectionId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid section ID"})
	}

	enrollments, err := h.enrollmentService.GetEnrollmentsByTeacherAndSection(userID, uint(sectionID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enrollments)
}

// GetMyEnrollments 获取学生自己的选课记录
func (h *EnrollmentHandler) GetMyEnrollments(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}

	enrollments, err := h.enrollmentService.GetEnrollmentsByStudentID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enrollments)
}

// EnrollCourse 学生选课
func (h *EnrollmentHandler) EnrollCourse(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}

	type EnrollRequest struct {
		SectionID uint `json:"sectionId"`
	}

	req := new(EnrollRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.enrollmentService.EnrollCourse(userID, req.SectionID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "选课成功"})
}

// DropCourse 学生退课
func (h *EnrollmentHandler) DropCourse(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	// 验证选课记录是否属于当前学生
	enrollment, err := h.enrollmentService.GetEnrollmentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "选课记录不存在"})
	}

	if enrollment.StudentID != userID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "只能退选自己的课程"})
	}

	if err := h.enrollmentService.DeleteEnrollment(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "退课成功"})
}

// UpdateEnrollment 更新选课记录（仅管理员）
func (h *EnrollmentHandler) UpdateEnrollment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	enrollment := new(model.Enrollment)
	if err := c.Bind(enrollment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	enrollment.ID = uint(id)
	if err := h.enrollmentService.UpdateEnrollment(enrollment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, enrollment)
}

// UpdateEnrollmentGrade 教师更新学生成绩
func (h *EnrollmentHandler) UpdateEnrollmentGrade(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	type GradeRequest struct {
		Grade string `json:"grade"`
	}

	req := new(GradeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 验证选课记录是否属于当前教师的课程
	enrollment, err := h.enrollmentService.GetEnrollmentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "选课记录不存在"})
	}

	// 检查该课程是否属于当前教师
	isTeacherCourse, err := h.enrollmentService.IsTeacherCourse(userID, enrollment.SectionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if !isTeacherCourse {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "只能为本人课程的学生更新成绩"})
	}

	enrollment.Grade = req.Grade
	if err := h.enrollmentService.UpdateEnrollment(enrollment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, enrollment)
}

// DeleteEnrollment 删除选课记录（仅管理员）
func (h *EnrollmentHandler) DeleteEnrollment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.enrollmentService.DeleteEnrollment(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
