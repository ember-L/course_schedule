package handler

import (
	"course_schedule/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SystemConfigHandler struct {
	service     service.SystemConfigService
	authService service.AuthService
}

func NewSystemConfigHandler(service service.SystemConfigService, authService service.AuthService) *SystemConfigHandler {
	return &SystemConfigHandler{service: service, authService: authService}
}

func (h *SystemConfigHandler) Register(e *echo.Echo) {
	authGroup := e.Group("/system")
	authGroup.Use(AuthMiddleware(h.authService))
	authGroup.GET("/config", h.Get)
	authGroup.PUT("/config", h.Update, RoleMiddleware("admin"))
}

func (h *SystemConfigHandler) Get(c echo.Context) error {
	cfg, err := h.service.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cfg)
}

func (h *SystemConfigHandler) Update(c echo.Context) error {
	var req struct {
		Semester  string `json:"semester"`
		StartDate string `json:"startDate"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	cfg, err := h.service.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	cfg.Semester = req.Semester
	cfg.StartDate = req.StartDate
	if err := h.service.Update(&cfg.SystemConfig); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cfg)
}
