@echo off
echo Starting development server with hot reloading...

REM 检查air是否已安装
where air >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo Air is not installed. Installing now...
    go install github.com/cosmtrek/air@latest
)

REM 切换到项目根目录
cd ..

REM 使用air启动应用
air