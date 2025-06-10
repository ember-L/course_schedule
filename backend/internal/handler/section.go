package handler

import (
	"course_schedule/internal/model"
	"course_schedule/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// SectionHandler 课程安排处理器
type SectionHandler struct {
	sectionService service.SectionService
}

// NewSectionHandler 创建课程安排处理器
func NewSectionHandler(sectionService service.SectionService) *SectionHandler {
	return &SectionHandler{
		sectionService: sectionService,
	}
}

// Register 注册路由
func (h *SectionHandler) Register(e *echo.Echo) {
	sections := e.Group("/sections")
	sections.GET("", h.GetAll)
	sections.GET("/:id", h.GetByID)
	sections.POST("", h.Create)
	sections.PUT("/:id", h.Update)
	sections.DELETE("/:id", h.Delete)
	sections.POST("/check-conflict", h.CheckConflict)
}

// GetAll 获取所有课程安排
func (h *SectionHandler) GetAll(c echo.Context) error {
	sections, err := h.sectionService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, sections)
}

// GetByID 通过ID获取课程安排
func (h *SectionHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的ID"})
	}

	section, err := h.sectionService.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if section == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "课程安排不存在"})
	}

	return c.JSON(http.StatusOK, section)
}

// Create 创建课程安排
func (h *SectionHandler) Create(c echo.Context) error {
	var section model.Section
	if err := c.Bind(&section); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求参数"})
	}

	if err := h.sectionService.Create(&section); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, section)
}

// Update 更新课程安排
func (h *SectionHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的ID"})
	}

	var section model.Section
	if err := c.Bind(&section); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求参数"})
	}

	section.ID = uint(id)
	if err := h.sectionService.Update(&section); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, section)
}

// Delete 删除课程安排
func (h *SectionHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的ID"})
	}

	if err := h.sectionService.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "课程安排已删除"})
}

// CheckConflict 检查课程安排冲突
func (h *SectionHandler) CheckConflict(c echo.Context) error {
	var section model.Section
	if err := c.Bind(&section); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求参数"})
	}

	hasConflict, message, err := h.sectionService.CheckConflict(&section)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"hasConflict": hasConflict,
		"message":     message,
	})
}