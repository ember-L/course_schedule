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
}

// NewEnrollmentHandler 创建选课记录处理器
func NewEnrollmentHandler(enrollmentService service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{enrollmentService: enrollmentService}
}

// Register 注册路由
func (h *EnrollmentHandler) Register(e *echo.Echo) {
	e.GET("/enrollments", h.GetAllEnrollments)
	e.GET("/enrollments/:id", h.GetEnrollmentByID)
	e.GET("/enrollments/student/:studentId", h.GetEnrollmentsByStudentID)
	e.GET("/enrollments/section/:sectionId", h.GetEnrollmentsBySectionID)
	e.POST("/enrollments", h.EnrollCourse)
	e.PUT("/enrollments/:id", h.UpdateEnrollment)
	e.DELETE("/enrollments/:id", h.DeleteEnrollment)
}

// GetAllEnrollments 获取所有选课记录
func (h *EnrollmentHandler) GetAllEnrollments(c echo.Context) error {
	enrollments, err := h.enrollmentService.GetAllEnrollments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enrollments)
}

// GetEnrollmentByID 根据ID获取选课记录
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

// GetEnrollmentsByStudentID 根据学生ID获取选课记录
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

// GetEnrollmentsBySectionID 根据课程安排ID获取选课记录
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

// EnrollCourse 学生选课
func (h *EnrollmentHandler) EnrollCourse(c echo.Context) error {
	type EnrollRequest struct {
		StudentID uint `json:"studentId"`
		SectionID uint `json:"sectionId"`
	}

	req := new(EnrollRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.enrollmentService.EnrollCourse(req.StudentID, req.SectionID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "选课成功"})
}

// UpdateEnrollment 更新选课记录
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

// DeleteEnrollment 删除选课记录
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
