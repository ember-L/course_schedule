package handler

import (
	"net/http"

	"course_schedule/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupMiddlewares 配置中间件
func SetupMiddlewares(e *echo.Echo) {
	// 日志中间件
	e.Use(middleware.Logger())

	// 恢复中间件
	e.Use(middleware.Recover())

	// CORS中间件
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
}

// AuthMiddleware 认证中间件
func AuthMiddleware(authService service.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 从请求头中获取令牌
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "未提供认证令牌"})
			}

			// 去掉Bearer前缀
			if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
				tokenString = tokenString[7:]
			}

			// 验证令牌
			claims, err := authService.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "无效的认证令牌"})
			}

			// 将用户信息存储到上下文中
			c.Set("userID", claims.UserID)
			c.Set("username", claims.Username)
			c.Set("userType", claims.UserType)

			return next(c)
		}
	}
}

// RoleMiddleware 角色权限中间件
func RoleMiddleware(requiredRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userType, ok := c.Get("userType").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型未找到"})
			}

			// 检查用户角色是否在允许的角色列表中
			hasRole := false
			for _, role := range requiredRoles {
				if userType == role {
					hasRole = true
					break
				}
			}

			if !hasRole {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "权限不足"})
			}

			return next(c)
		}
	}
}

// ResourceOwnerMiddleware 资源所有者权限中间件
func ResourceOwnerMiddleware(resourceType string, getOwnerID func(c echo.Context) (uint, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userType, ok := c.Get("userType").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型未找到"})
			}

			// 管理员可以访问所有资源
			if userType == "admin" {
				return next(c)
			}

			// 获取当前用户ID
			userID, ok := c.Get("userID").(uint)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户ID未找到"})
			}

			// 获取资源所有者ID
			ownerID, err := getOwnerID(c)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "无法获取资源信息"})
			}

			// 检查是否为资源所有者
			if userID != ownerID {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "只能访问自己的" + resourceType})
			}

			return next(c)
		}
	}
}

// APIKeyMiddleware API密钥中间件（用于外部API访问）
func APIKeyMiddleware(validAPIKeys []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get("X-API-Key")
			if apiKey == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "缺少API密钥"})
			}

			// 检查API密钥是否有效
			valid := false
			for _, key := range validAPIKeys {
				if apiKey == key {
					valid = true
					break
				}
			}

			if !valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "无效的API密钥"})
			}

			return next(c)
		}
	}
}

// RateLimitMiddleware 速率限制中间件
func RateLimitMiddleware(requestsPerMinute int) echo.MiddlewareFunc {
	// 简单的内存存储，生产环境建议使用Redis
	requestCounts := make(map[string]int)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP := c.RealIP()

			// 检查请求次数
			if requestCounts[clientIP] >= requestsPerMinute {
				return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "请求过于频繁，请稍后再试"})
			}

			requestCounts[clientIP]++

			return next(c)
		}
	}
}

// PermissionMiddleware 通用权限检查中间件
func PermissionMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			path := c.Request().URL.Path

			// 检查是否为公开API
			if CheckPermission("", method, path) {
				return next(c)
			}

			// 获取用户类型
			userType, ok := c.Get("userType").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型未找到"})
			}

			// 检查权限
			if !CheckPermission(userType, method, path) {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "权限不足"})
			}

			return next(c)
		}
	}
}

// 辅助函数
func getUserIDFromContext(c echo.Context) uint {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return 0
	}
	return userID
}

func getUserTypeFromContext(c echo.Context) string {
	userType, ok := c.Get("userType").(string)
	if !ok {
		return ""
	}
	return userType
}
