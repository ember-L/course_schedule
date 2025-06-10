#!/bin/bash

echo "Starting development server with hot reloading..."

# 检查air是否已安装
if ! command -v air &> /dev/null; then
    echo "Air is not installed. Installing now..."
    go install github.com/cosmtrek/air@latest
fi

# 切换到项目根目录
cd ..

# 使用air启动应用
air