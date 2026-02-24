#!/bin/bash

set -e

echo "🚀 SellCard 项目启动脚本"
echo "=========================="
echo ""

# 检查依赖
echo "📋 检查依赖..."

if ! command -v go &> /dev/null; then
    echo "❌ Go 未安装，请先安装 Go 1.20+"
    exit 1
fi

if ! command -v npm &> /dev/null; then
    echo "❌ Node.js/npm 未安装，请先安装 Node.js 16+"
    exit 1
fi

if ! command -v docker &> /dev/null; then
    echo "❌ Docker 未安装，请先安装 Docker"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose 未安装，请先安装 Docker Compose"
    exit 1
fi

echo "✅ Go 已安装: $(go version)"
echo "✅ npm 已安装: $(npm --version)"
echo "✅ Docker 已安装: $(docker --version)"
echo ""

# 启动 MySQL
echo "🗄️  启动 MySQL 服务..."
echo "   • 方式: Docker Compose"
echo "   • 端口: 3306"
echo "   • 密码: rootpwd"
echo ""

cd database
docker-compose up -d
cd ..

# 等待 MySQL 就绪
echo "⏳ 等待 MySQL 就绪..."
for i in {1..30}; do
    if docker exec database-db-1 mysql -uroot -prootpwd -e "SELECT 1" > /dev/null 2>&1; then
        echo "✅ MySQL 已就绪"
        break
    fi
    if [ $i -eq 30 ]; then
        echo "❌ MySQL 启动超时"
        exit 1
    fi
    echo "   等待中... ($i/30)"
    sleep 1
done
echo ""

# 启动后端
echo "🔧 启动后端服务器..."
echo "   • 配置文件: server/configs/config.yaml"
echo "   • Turnstile: 已禁用（开发模式）"
echo "   • 访问地址: http://localhost:8080"
echo ""

cd server
go run cmd/server/main.go &
BACKEND_PID=$!
echo "✅ 后端进程 ID: $BACKEND_PID"
echo ""

# 启动前端
echo "🎨 启动前端开发服务器..."
echo "   • 配置文件: web/.env.local"
echo "   • 访问地址: http://localhost:5173"
echo ""

cd ../web

# 检查 node_modules
if [ ! -d "node_modules" ]; then
    echo "📦 首次运行，安装依赖..."
    pnpm install
fi

pnpm dev &
FRONTEND_PID=$!
echo "✅ 前端进程 ID: $FRONTEND_PID"
echo ""

echo "=========================="
echo "🎉 项目已启动！"
echo ""
echo "访问地址："
echo "  • 前端: http://localhost:5173"
echo "  • 后端: http://localhost:8080"
echo "  • Swagger: http://localhost:8080/swagger/index.html"
echo ""
echo "数据库："
echo "  • 主机: localhost"
echo "  • 端口: 3306"
echo "  • 用户: root"
echo "  • 密码: rootpwd"
echo ""
echo "测试账号："
echo "  • 用户名: admin"
echo "  • 密码: 123456"
echo ""
echo "停止服务器："
echo "  按 Ctrl+C 停止前后端"
echo "  运行以下命令停止 MySQL："
echo "    cd database && docker-compose down"
echo ""

# 等待进程
wait
