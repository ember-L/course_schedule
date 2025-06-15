package handler

import (
	"net/http"
	"strconv"

	"course_schedule/internal/model"
	"course_schedule/internal/service"

	"github.com/labstack/echo/v4"
)

// TeacherHandler 教师处理器
type TeacherHandler struct {
	teacherService service.TeacherService
}

// NewTeacherHandler 创建教师处理器
func NewTeacherHandler(teacherService service.TeacherService) *TeacherHandler {
	return &TeacherHandler{teacherService: teacherService}
}

// Register 注册路由
func (h *TeacherHandler) Register(e *echo.Echo) {
	e.GET("/teachers", h.GetAllTeachers)
	e.GET("/teachers/:id", h.GetTeacherByID)
	e.POST("/teachers", h.CreateTeacher)
	e.PUT("/teachers/:id", h.UpdateTeacher)
	e.DELETE("/teachers/:id", h.DeleteTeacher)
}

// GetAllTeachers 获取所有教师
func (h *TeacherHandler) GetAllTeachers(c echo.Context) error {
	teachers, err := h.teacherService.GetAllTeachers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, teachers)
}

// GetTeacherByID 根据ID获取教师
func (h *TeacherHandler) GetTeacherByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	teacher, err := h.teacherService.GetTeacherByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Teacher not found"})
	}

	return c.JSON(http.StatusOK, teacher)
}

// CreateTeacher 创建教师
func (h *TeacherHandler) CreateTeacher(c echo.Context) error {
	teacher := new(model.Teacher)
	if err := c.Bind(teacher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.teacherService.CreateTeacher(teacher); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, teacher)
}

// UpdateTeacher 更新教师
func (h *TeacherHandler) UpdateTeacher(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	teacher := new(model.Teacher)
	if err := c.Bind(teacher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	teacher.ID = uint(id)
	if err := h.teacherService.UpdateTeacher(teacher); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, teacher)
}

// DeleteTeacher 删除教师
func (h *TeacherHandler) DeleteTeacher(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.teacherService.DeleteTeacher(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
