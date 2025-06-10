package handler

import (
	"net/http"
	"strconv"

	"course_schedule/internal/model"
	"course_schedule/internal/service"

	"github.com/labstack/echo/v4"
)

// StudentHandler 学生处理器
type StudentHandler struct {
	studentService service.StudentService
}

// NewStudentHandler 创建学生处理器
func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

// Register 注册路由
func (h *StudentHandler) Register(e *echo.Echo) {
	e.GET("/api/students", h.GetAllStudents)
	e.GET("/api/students/:id", h.GetStudentByID)
	e.GET("/api/students/studentId/:studentId", h.GetStudentByStudentID)
	e.POST("/api/students", h.CreateStudent)
	e.PUT("/api/students/:id", h.UpdateStudent)
	e.DELETE("/api/students/:id", h.DeleteStudent)
}

// GetAllStudents 获取所有学生
func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	students, err := h.studentService.GetAllStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, students)
}

// GetStudentByID 根据ID获取学生
func (h *StudentHandler) GetStudentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	student, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}

	return c.JSON(http.StatusOK, student)
}

// GetStudentByStudentID 根据学号获取学生
func (h *StudentHandler) GetStudentByStudentID(c echo.Context) error {
	studentID := c.Param("studentId")

	student, err := h.studentService.GetStudentByStudentID(studentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}

	return c.JSON(http.StatusOK, student)
}

// CreateStudent 创建学生
func (h *StudentHandler) CreateStudent(c echo.Context) error {
	student := new(model.Student)
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.studentService.CreateStudent(student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, student)
}

// UpdateStudent 更新学生
func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	student := new(model.Student)
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	student.ID = uint(id)
	if err := h.studentService.UpdateStudent(student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}

// DeleteStudent 删除学生
func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.studentService.DeleteStudent(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}