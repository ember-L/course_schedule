# API权限控制系统

## 概述

本系统实现了基于角色的访问控制（RBAC），支持三种用户角色：管理员（admin）、教师（teacher）和学生（student）。

## 权限级别

### 1. 公开API（Public）
无需认证即可访问的API：
- 课程信息查看
- 教室信息查看
- 时间段信息查看
- 用户登录

### 2. 管理员权限（Admin）
仅管理员可访问的API：
- 所有资源的增删改查
- 用户管理
- 系统配置

### 3. 教师权限（Teacher）
教师可访问的API：
- 查看个人信息
- 查看自己的课程
- 查看自己的课程安排
- 修改密码

### 4. 学生权限（Student）
学生可访问的API：
- 查看个人信息
- 查看自己的选课记录
- 选课/退课
- 修改密码

### 5. 管理员和教师权限（AdminTeacher）
管理员和教师都可以访问的API：
- 课程安排管理

### 6. 所有认证用户权限（AllAuth）
所有已认证用户都可以访问的API：
- 修改密码
- 查看课表

## 中间件系统

### 1. AuthMiddleware
认证中间件，验证JWT令牌并提取用户信息。

### 2. RoleMiddleware
角色中间件，检查用户是否具有指定角色。

### 3. PermissionMiddleware
通用权限中间件，基于权限配置进行动态检查。

### 4. ResourceOwnerMiddleware
资源所有者中间件，确保用户只能访问自己的资源。

### 5. APIKeyMiddleware
API密钥中间件，用于外部API访问。

### 6. RateLimitMiddleware
速率限制中间件，防止API滥用。

## 使用示例

### 路由注册示例

```go
// 公开路由
e.GET("/api/courses", h.GetAllCourses)

// 需要认证的路由
authGroup := e.Group("/api")
authGroup.Use(AuthMiddleware(authService))

// 管理员路由
adminGroup := authGroup.Group("/admin")
adminGroup.Use(RoleMiddleware("admin"))
adminGroup.GET("/profile/:id", h.GetAdminProfile)

// 多角色路由
adminTeacherGroup := authGroup.Group("/courses")
adminTeacherGroup.Use(RoleMiddleware("admin", "teacher"))
adminTeacherGroup.POST("", h.CreateCourse)
```

### 权限检查示例

```go
// 在处理器中检查资源所有者权限
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

    // 处理业务逻辑...
}
```

## 权限配置

权限配置在 `permissions.go` 文件中定义，支持以下配置：

- `Public`: 公开API列表
- `Admin`: 仅管理员可访问的API列表
- `Teacher`: 仅教师可访问的API列表
- `Student`: 仅学生可访问的API列表
- `AdminTeacher`: 管理员和教师可访问的API列表
- `AllAuth`: 所有认证用户可访问的API列表

## 安全特性

1. **JWT认证**: 使用JWT令牌进行用户认证
2. **角色验证**: 基于角色的权限控制
3. **资源隔离**: 用户只能访问自己的资源
4. **API密钥**: 支持外部API访问的密钥认证
5. **速率限制**: 防止API滥用
6. **CORS配置**: 跨域请求控制

## 错误处理

权限系统会返回以下HTTP状态码：

- `401 Unauthorized`: 未提供认证令牌或令牌无效
- `403 Forbidden`: 权限不足
- `429 Too Many Requests`: 请求过于频繁

## 最佳实践

1. **最小权限原则**: 只授予用户必要的最小权限
2. **资源隔离**: 确保用户只能访问自己的资源
3. **定期审查**: 定期审查和更新权限配置
4. **日志记录**: 记录所有权限相关的操作
5. **测试覆盖**: 为权限系统编写充分的测试用例 