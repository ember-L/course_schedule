package handler

import (
	"course_schedule/internal/model"
	"course_schedule/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ScheduleHandler 课表处理器
type ScheduleHandler struct {
	scheduleService service.ScheduleService
	authService     service.AuthService
}

// NewScheduleHandler 创建课表处理器
func NewScheduleHandler(scheduleService service.ScheduleService, authService service.AuthService) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: scheduleService,
		authService:     authService,
	}
}

// Register 注册路由
func (h *ScheduleHandler) Register(e *echo.Echo) {
	schedule := e.Group("/schedule")

	// 添加认证中间件
	schedule.Use(AuthMiddleware(h.authService))

	schedule.GET("/my", h.GetMySchedule)
	schedule.GET("/today", h.GetTodaySchedule)
	schedule.GET("/week", h.GetWeekSchedule)
	schedule.GET("/student/:studentId", h.GetStudentSchedule, RoleMiddleware("admin"))
	schedule.GET("/teacher/:teacherId", h.GetTeacherSchedule, RoleMiddleware("admin"))
	schedule.GET("/export", h.ExportSchedule)
}

// GetMySchedule 获取我的课表
func (h *ScheduleHandler) GetMySchedule(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}
	userType, ok := c.Get("userType").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型无效"})
	}
	sections, err := h.scheduleService.GetMySchedule(userID, userType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, sections)
}

// GetTodaySchedule 获取今日课表
func (h *ScheduleHandler) GetTodaySchedule(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}
	userType, ok := c.Get("userType").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型无效"})
	}
	sections, err := h.scheduleService.GetTodaySchedule(userID, userType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, sections)
}

// GetWeekSchedule 获取本周课表
func (h *ScheduleHandler) GetWeekSchedule(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}
	userType, ok := c.Get("userType").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型无效"})
	}
	sections, err := h.scheduleService.GetWeekSchedule(userID, userType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, sections)
}

// GetStudentSchedule 获取学生课表（管理员权限）
func (h *ScheduleHandler) GetStudentSchedule(c echo.Context) error {
	studentId, err := strconv.ParseUint(c.Param("studentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的学生ID"})
	}

	sections, err := h.scheduleService.GetStudentSchedule(uint(studentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, sections)
}

// GetTeacherSchedule 获取教师课表（管理员权限）
func (h *ScheduleHandler) GetTeacherSchedule(c echo.Context) error {
	teacherId, err := strconv.ParseUint(c.Param("teacherId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的教师ID"})
	}

	sections, err := h.scheduleService.GetTeacherSchedule(uint(teacherId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, sections)
}

// ExportSchedule 导出课表
func (h *ScheduleHandler) ExportSchedule(c echo.Context) error {
	// 从JWT中获取用户信息
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
	}

	userType, ok := c.Get("userType").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型无效"})
	}

	var sections []model.Section
	var err error

	switch userType {
	case "student":
		sections, err = h.scheduleService.GetStudentSchedule(userID)
	case "teacher":
		sections, err = h.scheduleService.GetTeacherSchedule(userID)
	case "admin":
		// 管理员可以导出指定用户的课表
		targetUserID := c.QueryParam("userId")
		targetUserType := c.QueryParam("userType")

		if targetUserID != "" && targetUserType != "" {
			uid, err := strconv.ParseUint(targetUserID, 10, 32)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的用户ID"})
			}

			switch targetUserType {
			case "student":
				sections, _ = h.scheduleService.GetStudentSchedule(uint(uid))
			case "teacher":
				sections, _ = h.scheduleService.GetTeacherSchedule(uint(uid))
			default:
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的用户类型"})
			}
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "缺少用户ID或用户类型参数"})
		}
	default:
		return c.JSON(http.StatusForbidden, map[string]string{"error": "权限不足"})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// TODO: 实现导出功能
	// 这里可以根据请求参数返回不同格式的数据（JSON、CSV等）
	// 暂时返回JSON格式
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":   sections,
		"format": "json",
		"total":  len(sections),
	})
}
