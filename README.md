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

## 开发环境设置

### 后端

1. 安装Go（1.16+）
2. 安装MySQL
3. 克隆仓库
4. 进入backend目录
5. 修改config/config.yaml中的数据库配置
6. 运行 `go mod tidy` 安装依赖
7. 运行 `go run cmd/app/main.go` 启动后端服务

#### 使用热重载（可选）

1. 安装Air: `go install github.com/cosmtrek/air@latest`
2. 在backend目录下运行: `air`

### 前端

1. 安装Node.js（14+）
2. 进入frontend-vite目录
3. 运行 `npm install` 安装依赖
4. 运行 `npm run dev` 启动开发服务器

## 部署

### 后端

1. 进入backend目录
2. 运行 `go build -o app cmd/app/main.go` 编译
3. 将编译后的二进制文件和config目录一起部署

### 前端

1. 进入frontend-vite目录
2. 运行 `npm run build` 构建
3. 将dist目录部署到Web服务器

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
