package handler

import (
	"course_schedule/internal/model"
	"course_schedule/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AnnouncementHandler struct {
	service     service.AnnouncementService
	authService service.AuthService
}

func NewAnnouncementHandler(service service.AnnouncementService, authService service.AuthService) *AnnouncementHandler {
	return &AnnouncementHandler{service: service, authService: authService}
}

func (h *AnnouncementHandler) Register(e *echo.Echo) {
	authGroup := e.Group("/announcements")
	authGroup.Use(AuthMiddleware(h.authService))
	authGroup.GET("", h.GetAll)
	authGroup.GET("/:id", h.GetByID)

	adminGroup := authGroup.Group("")
	adminGroup.Use(RoleMiddleware("admin"))
	adminGroup.POST("", h.Create)
	adminGroup.PUT("/:id", h.Update)
	adminGroup.DELETE("/:id", h.Delete)
}

func (h *AnnouncementHandler) GetAll(c echo.Context) error {
	list, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

func (h *AnnouncementHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID无效"})
	}
	a, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "公告不存在"})
	}
	return c.JSON(http.StatusOK, a)
}

func (h *AnnouncementHandler) Create(c echo.Context) error {
	a := new(model.Announcement)
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.service.Create(a); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, a)
}

func (h *AnnouncementHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID无效"})
	}
	a := new(model.Announcement)
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	a.ID = uint(id)
	if err := h.service.Update(a); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, a)
}

func (h *AnnouncementHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID无效"})
	}
	if err := h.service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
