package handler

import (
	"net/http"
	"strconv"

	"course_schedule/internal/model"
	"course_schedule/internal/service"

	"github.com/labstack/echo/v4"
)

// CourseHandler 课程处理器
type CourseHandler struct {
	courseService service.CourseService
}

// NewCourseHandler 创建课程处理器
func NewCourseHandler(courseService service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

// Register 注册路由
func (h *CourseHandler) Register(e *echo.Echo) {
	e.GET("/courses", h.GetAllCourses)
	e.GET("/courses/:id", h.GetCourseByID)
	e.POST("/courses", h.CreateCourse)
	e.PUT("/courses/:id", h.UpdateCourse)
	e.DELETE("/courses/:id", h.DeleteCourse)

	e.GET("/courses/teacher/:teacherId", h.GetCoursesByTeacherID)
}

// GetAllCourses 获取所有课程
func (h *CourseHandler) GetAllCourses(c echo.Context) error {
	courses, err := h.courseService.GetAllCourses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, courses)
}

// GetCourseByID 根据ID获取课程
func (h *CourseHandler) GetCourseByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	course, err := h.courseService.GetCourseByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Course not found"})
	}

	return c.JSON(http.StatusOK, course)
}

// CreateCourse 创建课程
func (h *CourseHandler) CreateCourse(c echo.Context) error {
	course := new(model.Course)
	if err := c.Bind(course); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.courseService.CreateCourse(course); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, course)
}

// UpdateCourse 更新课程
func (h *CourseHandler) UpdateCourse(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	course := new(model.Course)
	if err := c.Bind(course); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	course.ID = uint(id)
	if err := h.courseService.UpdateCourse(course); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, course)
}

// DeleteCourse 删除课程
func (h *CourseHandler) DeleteCourse(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.courseService.DeleteCourse(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// 根据教师ID获取课程
func (h *CourseHandler) GetCoursesByTeacherID(c echo.Context) error {

	teacherID, err := strconv.ParseUint(c.Param("teacherId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid teacher ID"})
	}
	courses, err := h.courseService.GetByTeacherID(uint(teacherID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, courses)
}
