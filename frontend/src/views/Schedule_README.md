# 课表系统重构说明

## 概述

本次重构实现了基于角色的课表查看功能，支持不同用户类型查看相应的课表信息。

## 功能特性

### 1. 角色权限控制

#### 学生用户
- 查看自己的课表
- 按学期筛选课程
- 导出个人课表（PDF/Excel）
- 查看课表统计信息

#### 教师用户
- 查看自己的授课安排
- 按学期筛选课程
- 导出个人课表（PDF/Excel）
- 查看课表统计信息

#### 管理员用户
- 查看所有学生的课表
- 查看所有教师的课表
- 按用户类型和用户筛选
- 导出任意用户的课表
- 查看课表统计信息

### 2. 课表统计功能

- **总课程数**: 当前筛选条件下的课程总数
- **本周课程**: 本周内的课程数量
- **今日课程**: 今天的课程数量
- **总学时**: 所有课程的总学时

### 3. 导出功能

- **PDF导出**: 生成可打印的PDF格式课表
- **Excel导出**: 生成可编辑的Excel格式课表
- 支持按学期导出
- 文件名包含用户信息和学期信息

## 技术实现

### 前端架构

#### 组件结构
```
Schedule.vue (主组件)
├── TimeTable.vue (课表显示组件)
├── 用户选择器 (管理员)
├── 学期选择器
├── 导出功能
└── 统计信息面板
```

#### 状态管理
- 使用Vue 3 Composition API
- 集成认证状态管理
- 响应式数据绑定

#### API集成
```javascript
// 新增的API接口
export const scheduleApi = {
  getMySchedule: () => api.get('/schedule/my'),
  getTodaySchedule: () => api.get('/schedule/today'),
  getWeekSchedule: () => api.get('/schedule/week'),
  getStudentSchedule: (studentId) => api.get(`/schedule/student/${studentId}`),
  getTeacherSchedule: (teacherId) => api.get(`/schedule/teacher/${teacherId}`),
  exportSchedule: (params) => api.get('/schedule/export', { params }),
};
```

### 后端架构

#### 新增API端点
```
GET /schedule/my          - 获取我的课表
GET /schedule/today       - 获取今日课表
GET /schedule/week        - 获取本周课表
GET /schedule/student/:id - 获取学生课表（管理员）
GET /schedule/teacher/:id - 获取教师课表（管理员）
GET /schedule/export      - 导出课表
GET /sections/teacher/:id - 获取教师课程安排
GET /sections/student/:id - 获取学生课程安排
```

#### 权限控制
- JWT认证中间件
- 角色权限验证
- 资源所有者验证

#### 数据查询优化
- 预加载关联数据（Course, Teacher, Classroom, TimeSlot）
- 通过JOIN查询学生选课记录
- 按教师ID直接查询课程安排

## 使用说明

### 学生用户
1. 登录后自动显示个人课表
2. 选择学期查看特定学期的课程
3. 点击导出按钮下载课表

### 教师用户
1. 登录后自动显示个人授课安排
2. 选择学期查看特定学期的课程
3. 点击导出按钮下载课表

### 管理员用户
1. 在用户选择器中选择要查看的用户
2. 选择学期查看特定学期的课程
3. 点击导出按钮下载课表

## 数据流程

### 学生课表查询流程
1. 前端调用 `/sections/student/:id` API
2. 后端通过JOIN查询获取学生选课记录
3. 预加载课程、教师、教室、时间段信息
4. 返回完整的课表数据

### 教师课表查询流程
1. 前端调用 `/sections/teacher/:id` API
2. 后端直接查询教师的课程安排
3. 预加载课程、教师、教室、时间段信息
4. 返回完整的课表数据

### 管理员查询流程
1. 管理员选择用户类型和具体用户
2. 根据用户类型调用相应的API
3. 后端验证管理员权限
4. 返回指定用户的课表数据

## 权限配置

### 前端权限控制
```javascript
const isAdmin = computed(() => authStore.isAdmin);
const isTeacher = computed(() => authStore.isTeacher);
const isStudent = computed(() => authStore.isStudent);
```

### 后端权限控制
```go
// 检查管理员权限
userType := c.Get("userType").(string)
if userType != "admin" {
    return c.JSON(http.StatusForbidden, map[string]string{"error": "权限不足"})
}
```

## 样式设计

### 响应式布局
- 桌面端：完整功能展示
- 平板端：适配中等屏幕
- 移动端：简化界面，优化触摸操作

### 主题适配
- 支持浅色/深色主题
- 使用CSS变量实现主题切换
- 导出时保持样式一致性

## 性能优化

### 前端优化
- 使用计算属性缓存筛选结果
- 懒加载课表组件
- 防抖处理用户输入

### 后端优化
- 数据库查询优化
- 预加载关联数据
- 缓存常用查询结果

## 错误处理

### 前端错误处理
- API请求错误提示
- 网络异常处理
- 用户友好的错误信息

### 后端错误处理
- 参数验证错误
- 数据库查询错误
- 权限验证错误

## 测试建议

### 功能测试
1. 不同角色用户登录测试
2. 课表数据正确性测试
3. 导出功能测试
4. 权限控制测试

### 性能测试
1. 大量数据加载测试
2. 并发访问测试
3. 导出功能性能测试

### 兼容性测试
1. 不同浏览器兼容性
2. 移动端适配测试
3. 不同屏幕尺寸测试

## 部署说明

### 前端部署
1. 构建生产版本
2. 配置API基础URL
3. 部署到Web服务器

### 后端部署
1. 编译Go程序
2. 配置数据库连接
3. 设置环境变量
4. 启动服务

## 维护说明

### 日常维护
1. 监控API响应时间
2. 检查数据库性能
3. 更新依赖包版本

### 故障排查
1. 查看应用日志
2. 检查数据库连接
3. 验证权限配置
4. 测试API端点

## 扩展建议

### 功能扩展
1. 添加课程冲突检测
2. 实现课表分享功能
3. 支持自定义时间段
4. 添加课程提醒功能

### 技术扩展
1. 实现WebSocket实时更新
2. 添加离线缓存功能
3. 支持PWA应用
4. 集成第三方日历

## 总结

本次重构成功实现了基于角色的课表查看功能，提供了良好的用户体验和完整的权限控制。系统具有良好的可扩展性和维护性，为后续功能开发奠定了坚实的基础。 