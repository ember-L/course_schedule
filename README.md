# 高校排课系统

一个使用Echo+MySQL+Vue3(ArcoDesign)+Viper的高校排课系统。

## 项目结构

```
course_schedule/
├── backend/             # Echo + MySQL + Viper 后端
│   ├── cmd/             # 应用程序入口
│   ├── config/          # 配置文件
│   ├── internal/        # 内部包
│   │   ├── handler/     # HTTP 处理器
│   │   ├── model/       # 数据模型
│   │   ├── repository/  # 数据库操作
│   │   └── service/     # 业务逻辑
│   ├── pkg/             # 可导出的包
│   └── scripts/         # 数据库脚本等
└── frontend/            # Vue3 + ArcoDesign 前端
    ├── public/          # 静态资源
    └── src/             # 源代码
        ├── api/         # API 调用
        ├── assets/      # 资源文件
        ├── components/  # 组件
        ├── router/      # 路由
        ├── store/       # 状态管理
        └── views/       # 页面视图
```



## 快速开始

克隆仓库

```sh
git clone https://github.com/ember-L/course_schedule.git
```



### 后端

1. 安装Go（1.16+）和MySQL
2. 克隆仓库：
3. 配置数据库
```sql
CREATE DATABASE course_schedule CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

4. 修改配置文件 `backend/config/config.yaml`
5. 运行后端服务

```sh
cd backend
go mod tidy
go run cmd/app/main.go
```

#### 使用热重载（可选）

1. 安装Air: `go install github.com/cosmtrek/air@latest`
2. 在backend目录下运行: `air`

### 前端

1. 安装Node.js（16+）
2. 安装依赖并启动开发服务器

```bash
cd frontend
npm install
npm run dev
```

## 功能特性

- 课程管理：添加、编辑、删除课程信息
- 教师管理：添加、编辑、删除教师信息
- 教室管理：添加、编辑、删除教室信息
- 排课管理：安排课程、教师、教室、时间段，自动检测冲突
- 学生选课：学生可以选择课程，查看自己的课表

## 技术栈

### 后端

- Echo：Go语言高性能Web框架
- MySQL：关系型数据库
- Viper：配置管理
- GORM：ORM库
- Air：热重载工具

### 前端

- Vue3：渐进式JavaScript框架
- Vite：现代前端构建工具，提供更快的开发体验
- ArcoDesign：字节跳动开源的企业级设计系统
- Vue Router：路由管理
- Pinia：状态管理
- Axios：HTTP客户端



## 开发指南

### 后端开发

1. 使用Air实现热重载

```bash
cd backend
air
```

1. 添加新的API端点在 `internal/model`中定义模型

2. 在 `internal/repository`中实现数据访问

3. 在 `internal/service`中实现业务逻辑

4. 在 `internal/handler`中实现HTTP处理器

   

### 前端开发

1. 添加新的页面
2. 在 `src/views`中创建Vue组件
3. 在 `src/router/index.js`中添加路由
4. 在 `src/api`中添加API调用
5. 构建生产版本

```bash
npm run build
```



## 数据库设计

系统包含以下主要数据表：

- teachers：教师信息
- courses：课程信息
- classrooms：教室信息
- time_slots：时间段
- sections：课程安排
- students：学生信息
- enrollments：选课记录

## 贡献

欢迎提交问题和功能请求！
