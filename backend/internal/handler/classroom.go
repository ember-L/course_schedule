package handler

import (
	"net/http"
	"strconv"

	"course_schedule/internal/model"
	"course_schedule/internal/service"

	"github.com/labstack/echo/v4"
)

// ClassroomHandler 教室处理器
type ClassroomHandler struct {
	classroomService service.ClassroomService
	authService      service.AuthService
}

// NewClassroomHandler 创建教室处理器
func NewClassroomHandler(classroomService service.ClassroomService, authService service.AuthService) *ClassroomHandler {
	return &ClassroomHandler{
		classroomService: classroomService,
		authService:      authService,
	}
}

// Register 注册路由
func (h *ClassroomHandler) Register(e *echo.Echo) {
	// 公开路由 - 所有用户都可以查看教室信息
	e.GET("/classrooms", h.GetAllClassrooms)
	e.GET("/classrooms/:id", h.GetClassroomByID)

	// 需要认证的路由
	authGroup := e.Group("")
	authGroup.Use(AuthMiddleware(h.authService))

	// 管理员可以管理的路由
	adminGroup := authGroup.Group("/classrooms")
	adminGroup.Use(RoleMiddleware("admin"))
	adminGroup.POST("", h.CreateClassroom)
	adminGroup.PUT("/:id", h.UpdateClassroom)
	adminGroup.DELETE("/:id", h.DeleteClassroom)
}

// GetAllClassrooms 获取所有教室
func (h *ClassroomHandler) GetAllClassrooms(c echo.Context) error {
	classrooms, err := h.classroomService.GetAllClassrooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, classrooms)
}

// GetClassroomByID 根据ID获取教室
func (h *ClassroomHandler) GetClassroomByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	classroom, err := h.classroomService.GetClassroomByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Classroom not found"})
	}

	return c.JSON(http.StatusOK, classroom)
}

// CreateClassroom 创建教室
func (h *ClassroomHandler) CreateClassroom(c echo.Context) error {
	classroom := new(model.Classroom)
	if err := c.Bind(classroom); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.classroomService.CreateClassroom(classroom); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, classroom)
}

// UpdateClassroom 更新教室
func (h *ClassroomHandler) UpdateClassroom(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	classroom := new(model.Classroom)
	if err := c.Bind(classroom); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	classroom.ID = uint(id)
	if err := h.classroomService.UpdateClassroom(classroom); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, classroom)
}

// DeleteClassroom 删除教室
func (h *ClassroomHandler) DeleteClassroom(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.classroomService.DeleteClassroom(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
