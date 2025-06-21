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
	authService   service.AuthService
}

// NewCourseHandler 创建课程处理器
func NewCourseHandler(courseService service.CourseService, authService service.AuthService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
		authService:   authService,
	}
}

// Register 注册路由
func (h *CourseHandler) Register(e *echo.Echo) {
	// 公开路由 - 所有用户都可以查看课程信息
	e.GET("/courses", h.GetAllCourses)
	e.GET("/courses/:id", h.GetCourseByID)

	// 需要认证的路由
	authGroup := e.Group("")
	authGroup.Use(AuthMiddleware(h.authService))

	// 管理员和教师可以管理的路由
	adminTeacherGroup := authGroup.Group("/courses")
	adminTeacherGroup.Use(RoleMiddleware("admin", "teacher"))
	adminTeacherGroup.POST("", h.CreateCourse)
	adminTeacherGroup.PUT("/:id", h.UpdateCourse)
	adminTeacherGroup.DELETE("/:id", h.DeleteCourse)

	// 教师可以查看自己的课程
	teacherGroup := authGroup.Group("/courses")
	teacherGroup.Use(RoleMiddleware("teacher"))
	teacherGroup.GET("/teacher/:teacherId", h.GetCoursesByTeacherID)
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

	// 检查权限：教师只能查看自己的课程
	currentUserID := getUserIDFromContext(c)
	currentUserType := getUserTypeFromContext(c)

	if currentUserType == "teacher" && currentUserID != uint(teacherID) {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "只能查看自己的课程"})
	}

	courses, err := h.courseService.GetByTeacherID(uint(teacherID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, courses)
}
