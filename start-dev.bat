@echo off
setlocal enabledelayedexpansion

echo.
echo 🚀 SellCard 项目启动脚本
echo ==========================
echo.

REM 检查 Go
go version >nul 2>&1
if errorlevel 1 (
    echo ❌ Go 未安装，请先安装 Go 1.20+
    exit /b 1
)

REM 检查 npm
npm --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Node.js/npm 未安装，请先安装 Node.js 16+
    exit /b 1
)

REM 检查 Docker
docker --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Docker 未安装，请先安装 Docker
    exit /b 1
)

REM 检查 Docker Compose
docker-compose --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Docker Compose 未安装，请先安装 Docker Compose
    exit /b 1
)

echo 📋 检查依赖...
echo ✅ Go 已安装
for /f "tokens=*" %%i in ('npm --version') do set NPM_VERSION=%%i
echo ✅ npm 已安装: !NPM_VERSION!
for /f "tokens=*" %%i in ('docker --version') do set DOCKER_VERSION=%%i
echo ✅ Docker 已安装: !DOCKER_VERSION!
echo.

REM 启动 MySQL
echo 🗄️  启动 MySQL 服务...
echo    • 方式: Docker Compose
echo    • 端口: 3306
echo    • 密码: rootpwd
echo.

cd database
call docker-compose up -d
cd ..

REM 等待 MySQL 就绪
echo ⏳ 等待 MySQL 就绪...
setlocal enabledelayedexpansion
for /L %%i in (1,1,30) do (
    docker exec database-db-1 mysql -uroot -prootpwd -e "SELECT 1" >nul 2>&1
    if !errorlevel! equ 0 (
        echo ✅ MySQL 已就绪
        goto mysql_ready
    )
    if %%i lss 30 (
        echo    等待中... (%%i/30)
        timeout /t 1 /nobreak >nul
    )
)
echo ❌ MySQL 启动超时
exit /b 1

:mysql_ready
echo.

REM 启动后端
echo 🔧 启动后端服务器...
echo    • 配置文件: server/configs/config.yaml
echo    • Turnstile: 已禁用（开发模式）
echo    • 访问地址: http://localhost:8080
echo.

cd server
start "SellCard Backend" cmd /k "go run cmd/server/main.go"
cd ..

timeout /t 2 /nobreak
echo ✅ 后端已启动
echo.

REM 启动前端
echo 🎨 启动前端开发服务器...
echo    • 配置文件: web/.env.local
echo    • 访问地址: http://localhost:5173
echo.

cd web

REM 检查 node_modules
if not exist "node_modules" (
    echo 📦 首次运行，安装依赖...
    call pnpm install
)

start "SellCard Frontend" cmd /k "pnpm dev"
cd ..

echo ✅ 前端已启动
echo.
echo ==========================
echo 🎉 项目已启动！
echo.
echo 访问地址：
echo   • 前端: http://localhost:5173
echo   • 后端: http://localhost:8080
echo   • Swagger: http://localhost:8080/swagger/index.html
echo.
echo 数据库：
echo   • 主机: localhost
echo   • 端口: 3306
echo   • 用户: root
echo   • 密码: rootpwd
echo.
echo 测试账号：
echo   • 用户名: admin
echo   • 密码: 123456
echo.
echo 停止服务器：
echo   按 Ctrl+C 停止各个服务器窗口
echo   运行以下命令停止 MySQL：
echo     cd database ^& docker-compose down
echo.
