## Docker部署

本项目支持使用Docker进行容器化部署，提供了完整的Docker配置文件。

### 目录结构

```sh
docker/
├── backend/ # 后端Docker配置
│ └── Dockerfile
├── frontend/ # 前端Docker配置
│ ├── Dockerfile
│ └── nginx.conf
└── docker-compose.yml # Docker Compose配置
```



### 后端Dockerfile

```dockerfile
# 构建阶段
FROM golang:1.18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/app/main.go

# 运行阶段
FROM alpine:latest

# 安装必要的工具
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/app .
COPY --from=builder /app/config ./config

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./app"]
```



### 前端dockfile

```dockerfile
# 构建阶段
FROM node:16-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制源代码
COPY . .

# 构建应用
RUN npm run build

# 运行阶段
FROM nginx:alpine

# 复制构建产物到Nginx目录
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制Nginx配置
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

# 启动Nginx
CMD ["nginx", "-g", "daemon off;"]
```



### nginx配置

```nginx
server {
    listen 80;
    server_name localhost;
    
    location / {
        root /usr/share/nginx/html;
        index index.html;
        try_files $uri $uri/ /index.html;
    }
    
    location /api/ {
        proxy_pass http://backend:8080/api/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}

```



### Docker Compose

```dockerfile
version: '3'

services:
  mysql:
    image: mysql:5.7
    container_name: course_schedule_mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: course_schedule
      MYSQL_USER: user
      MYSQL_PASSWORD: password
    volumes:
      - mysql_data:/var/lib/mysql
      - ./scripts/test_data.sql:/docker-entrypoint-initdb.d/test_data.sql
    ports:
      - "3306:3306"
    networks:
      - app_network

  backend:
    build: 
      context: ./backend
      dockerfile: Dockerfile
    container_name: course_schedule_backend
    restart: always
    depends_on:
      - mysql
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USERNAME: user
      DB_PASSWORD: password
      DB_NAME: course_schedule
    ports:
      - "8080:8080"
    networks:
      - app_network

  frontend:
    build:
      context: ./frontend-vite
      dockerfile: Dockerfile
    container_name: course_schedule_frontend
    restart: always
    depends_on:
      - backend
    ports:
      - "80:80"
    networks:
      - app_network

networks:
  app_network:
    driver: bridge

volumes:
  mysql_data:

```

### 部署步骤

1. 确保已安装Docker和Docker Compose

   ```bash
   docker --version
   docker-compose --version
   ```

2. 克隆仓库并进入项目目录

   ```bash
   git clone https://github.com/yourusername/course_schedule.git
   cd course_schedule
   ```

   

3. 启动服务

   ```bash
   docker-compose up -d
   ```

   

4. 访问应用

   - 前端: http://localhost
   - 后端API: http://localhost:8080/api

5. 查看日志

   ```bash
   # 查看所有容器日志
   docker-compose logs
   
   # 查看特定服务日志
   docker-compose logs backend
   docker-compose logs frontend
   ```

   

6. 停止服务

   ```bash
   docker-compose down
   ```

