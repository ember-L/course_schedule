package handler

import (
	"net/http"
	"strconv"

	"course_schedule/internal/model"
	"course_schedule/internal/service"

	"github.com/labstack/echo/v4"
)

// TimeSlotHandler 时间段处理器
type TimeSlotHandler struct {
	timeSlotService service.TimeSlotService
}

// NewTimeSlotHandler 创建时间段处理器
func NewTimeSlotHandler(timeSlotService service.TimeSlotService) *TimeSlotHandler {
	return &TimeSlotHandler{timeSlotService: timeSlotService}
}

// Register 注册路由
func (h *TimeSlotHandler) Register(e *echo.Echo) {
	e.GET("/timeSlots", h.GetAllTimeSlots)
	e.GET("/timeSlots/:id", h.GetTimeSlotByID)
	e.POST("/timeSlots", h.CreateTimeSlot)
	e.PUT("/timeSlots/:id", h.UpdateTimeSlot)
	e.DELETE("/timeSlots/:id", h.DeleteTimeSlot)
}

// GetAllTimeSlots 获取所有时间段
func (h *TimeSlotHandler) GetAllTimeSlots(c echo.Context) error {
	timeSlots, err := h.timeSlotService.GetAllTimeSlots()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, timeSlots)
}

// GetTimeSlotByID 根据ID获取时间段
func (h *TimeSlotHandler) GetTimeSlotByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	timeSlot, err := h.timeSlotService.GetTimeSlotByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "TimeSlot not found"})
	}

	return c.JSON(http.StatusOK, timeSlot)
}

// CreateTimeSlot 创建时间段
func (h *TimeSlotHandler) CreateTimeSlot(c echo.Context) error {
	timeSlot := new(model.TimeSlot)
	if err := c.Bind(timeSlot); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.timeSlotService.CreateTimeSlot(timeSlot); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, timeSlot)
}

// UpdateTimeSlot 更新时间段
func (h *TimeSlotHandler) UpdateTimeSlot(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	timeSlot := new(model.TimeSlot)
	if err := c.Bind(timeSlot); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	timeSlot.ID = uint(id)
	if err := h.timeSlotService.UpdateTimeSlot(timeSlot); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, timeSlot)
}

// DeleteTimeSlot 删除时间段
func (h *TimeSlotHandler) DeleteTimeSlot(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.timeSlotService.DeleteTimeSlot(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
