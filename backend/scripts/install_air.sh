#!/bin/bash

# 安装Air热重载工具
echo "Installing Air for hot reloading..."
go install github.com/cosmtrek/air@latest

echo "Air installed successfully!"
echo "Run 'air' in your project root to start hot reloading."