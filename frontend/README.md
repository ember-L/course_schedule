# 高校排课系统前端

基于Vue3、Vite和ArcoDesign的高校排课系统前端。

## 技术栈

- Vue3：渐进式JavaScript框架
- Vite：现代前端构建工具
- ArcoDesign：字节跳动开源的企业级设计系统
- Vue Router：路由管理
- Pinia：状态管理
- Axios：HTTP客户端

## 项目结构

```
frontend/
├── public/          # 静态资源
└── src/             # 源代码
    ├── api/         # API 调用
    ├── assets/      # 资源文件
    ├── components/  # 组件
    ├── router/      # 路由
    ├── store/       # 状态管理
    └── views/       # 页面视图
```

## 开发环境设置

1. 安装依赖
```bash
npm install
```

2. 启动开发服务器
```bash
npm run dev
```

3. 构建生产版本
```bash
npm run build
```

## 功能模块

- 课程管理：添加、编辑、删除课程信息
- 教师管理：添加、编辑、删除教师信息
- 教室管理：添加、编辑、删除教室信息
- 排课管理：安排课程、教师、教室、时间段，自动检测冲突
- 学生管理：添加、编辑、删除学生信息
- 选课管理：学生选课、查看课表