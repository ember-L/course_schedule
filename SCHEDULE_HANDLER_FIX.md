# Schedule Handler 修复说明

## 问题概述

在 `backend/internal/handler/schedule.go` 文件中发现了多个问题需要修复：

1. **类型不匹配**: `sectionService.GetByStudent` 和 `GetByTeacher` 返回 `[]model.Section`，但代码中声明为 `[]interface{}`
2. **缺少中间件**: 路由没有添加认证中间件
3. **错误处理**: 需要更好的错误处理机制
4. **权限验证**: 需要更完善的权限验证
5. **函数重复定义**: `CheckPermission` 函数重复定义

## 修复内容

### 1. 修复类型不匹配问题

**问题**: 
```go
var sections []interface{}  // 错误的类型声明
```

**修复**:
```go
var sections []model.Section  // 正确的类型声明
```

### 2. 添加认证中间件

**问题**: 路由没有认证保护

**修复**:
```go
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
```

### 3. 改进错误处理

**问题**: 类型断言没有错误检查

**修复**:
```go
// 从JWT中获取用户信息
userID, ok := c.Get("userID").(uint)
if !ok {
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户信息无效"})
}

userType, ok := c.Get("userType").(string)
if !ok {
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户类型无效"})
}
```

### 4. 使用 switch 语句优化逻辑

**问题**: 使用 if-else 链，代码冗长

**修复**:
```go
switch userType {
case "student":
	sections, err = h.sectionService.GetByStudent(userID)
case "teacher":
	sections, err = h.sectionService.GetByTeacher(userID)
default:
	return c.JSON(http.StatusForbidden, map[string]string{"error": "权限不足"})
}
```

### 5. 添加 authService 依赖

**问题**: `AuthMiddleware` 需要 `authService` 参数

**修复**:
```go
// ScheduleHandler 课表处理器
type ScheduleHandler struct {
	sectionService    service.SectionService
	enrollmentService service.EnrollmentService
	authService       service.AuthService
}

// NewScheduleHandler 创建课表处理器
func NewScheduleHandler(sectionService service.SectionService, enrollmentService service.EnrollmentService, authService service.AuthService) *ScheduleHandler {
	return &ScheduleHandler{
		sectionService:    sectionService,
		enrollmentService: enrollmentService,
		authService:       authService,
	}
}
```

### 6. 更新主应用文件

**问题**: 主应用中没有传递 `authService` 参数

**修复**:
```go
// backend/cmd/app/main.go
scheduleHandler := handler.NewScheduleHandler(sectionService, enrollmentService, authService)
```

### 7. 改进导出功能

**问题**: 导出功能返回格式不统一

**修复**:
```go
// 暂时返回JSON格式
return c.JSON(http.StatusOK, map[string]interface{}{
	"data": sections,
	"format": "json",
	"total": len(sections),
})
```

## 权限控制

### 1. 认证中间件
- 验证JWT令牌
- 提取用户信息到上下文
- 处理无效令牌

### 2. 角色中间件
- 检查用户角色权限
- 支持多个角色验证
- 提供清晰的错误消息

### 3. 权限配置
- 使用 `permissions.go` 中的权限配置
- 支持公开API、角色特定API
- 灵活的权限管理

## API端点说明

### 1. 我的课表
```
GET /schedule/my
```
- 需要认证
- 学生和教师可访问
- 返回当前用户的课表

### 2. 今日课表
```
GET /schedule/today
```
- 需要认证
- 学生和教师可访问
- 返回今日课程（待实现过滤逻辑）

### 3. 本周课表
```
GET /schedule/week
```
- 需要认证
- 学生和教师可访问
- 返回本周课程（待实现过滤逻辑）

### 4. 学生课表（管理员）
```
GET /schedule/student/:studentId
```
- 需要管理员权限
- 返回指定学生的课表

### 5. 教师课表（管理员）
```
GET /schedule/teacher/:teacherId
```
- 需要管理员权限
- 返回指定教师的课表

### 6. 导出课表
```
GET /schedule/export
```
- 需要认证
- 支持所有用户类型
- 管理员可导出指定用户的课表

## 错误处理

### 1. 认证错误
- 401 Unauthorized: 未提供令牌或令牌无效
- 清晰的错误消息

### 2. 权限错误
- 403 Forbidden: 权限不足
- 角色验证失败

### 3. 参数错误
- 400 Bad Request: 无效的用户ID或类型
- 参数验证失败

### 4. 服务器错误
- 500 Internal Server Error: 数据库查询失败
- 详细的错误信息

## 测试建议

### 1. 认证测试
- [ ] 无令牌访问API
- [ ] 无效令牌访问API
- [ ] 有效令牌访问API

### 2. 权限测试
- [ ] 学生访问教师API
- [ ] 教师访问管理员API
- [ ] 管理员访问所有API

### 3. 功能测试
- [ ] 学生查看自己的课表
- [ ] 教师查看自己的课表
- [ ] 管理员查看任意用户课表
- [ ] 导出功能测试

### 4. 错误处理测试
- [ ] 无效用户ID
- [ ] 无效用户类型
- [ ] 数据库连接失败
- [ ] 网络异常

## 部署注意事项

### 1. 环境变量
- 确保JWT密钥配置正确
- 数据库连接配置
- CORS配置

### 2. 中间件顺序
- 认证中间件必须在权限中间件之前
- 错误处理中间件应该在最后

### 3. 日志记录
- 记录认证失败
- 记录权限验证失败
- 记录API调用统计

## 总结

通过以上修复，Schedule Handler 现在具备了：

1. **完整的认证机制**: JWT令牌验证
2. **完善的权限控制**: 基于角色的访问控制
3. **健壮的错误处理**: 详细的错误信息和状态码
4. **类型安全**: 正确的Go类型声明
5. **代码优化**: 使用switch语句和更好的结构

这些修复确保了课表API的安全性、可靠性和可维护性。 