# Sellcard — 后台管理示例（Gin + Vue 3 + Vuetify）

本仓库包含一个后端（Gin + GORM + MySQL）和一个前端（Vite + Vue 3 + Pinia + Vuetify）示例骨架，演示登录到后台管理的完整流程（前后端分离）。

## 系统要求

| 工具 | 版本 | 用途 |
|------|------|------|
| **Go** | ≥ 1.20 | 后端编译运行 |
| **Node.js/npm** | ≥ 16 | 前端编译运行 |
| **Docker** | Latest | 运行 MySQL 容器 |
| **Docker Compose** | ≥ 1.29 | 编排 MySQL 容器 |

> ✅ 推荐：使用 [Docker Desktop](https://www.docker.com/products/docker-desktop) 一键安装 Docker + Compose

## 项目结构

### 后端（server）
- 入口: `/server/cmd/server/main.go`
- 配置: `/server/configs/config.yaml`（已忽略，使用环境变量）
- 数据模型: `/server/internal/model`
- 仓库/服务/控制器: `/server/internal`
- 数据库初始化 SQL: `/server/scripts/db/init.sql`
- **新增**: Turnstile 人机验证集成

### 前端（web）
- 基于: Vite + Vue 3 + Pinia + Vuetify
- 入口: `/web/src/main.ts`
- 登录页: `/web/src/pages/login.vue`
- 后台管理: `/web/src/pages/backend.vue`
- **新增**: Turnstile 验证集成

## 快速启动

### 一键启动（推荐）

**Linux/macOS：**
```bash
chmod +x start-dev.sh
./start-dev.sh
```

**Windows：**
```bash
start-dev.bat
```

> 📌 **要求**：已安装 Docker 和 Docker Compose（用于启动 MySQL）

该脚本会自动启动：
1. ✅ MySQL 数据库（Docker）
2. ✅ 后端服务器（Go）
3. ✅ 前端开发服务器（Node.js）

### 手动启动

**Step 1：启动 MySQL**
```bash
cd database
docker-compose up -d
cd ..
```

**Step 2：启动后端**
```bash
cd server
go mod tidy
go run cmd/server/main.go
```
- 访问地址: `http://localhost:8080`
- Swagger 文档: `http://localhost:8080/swagger/index.html`

**Step 3：启动前端**（在新终端）
```bash
cd web
npm install
npm run dev
```
- 访问地址: `http://localhost:5173`

## 测试账号

```
用户名: admin
密码: 123456
```

## 环境配置

### 本地开发（零配置）

项目已配置为开发默认值：
- ✅ 后端 Turnstile **禁用**
- ✅ 前端 `.env.local` 已配置
- ✅ 本地 MySQL 默认配置
- ✅ 开箱即用，无需修改配置

### 生产部署

详见 **[ENV_SETUP.md](ENV_SETUP.md)** 获取以下内容：
- 📝 环境变量列表
- 🐳 Docker 部署指南
- 🔐 敏感信息安全管理
- 🚀 CI/CD 集成示例

### 支持的环境变量

```bash
# Turnstile（人机验证）
TURNSTILE_ENABLED=true
TURNSTILE_SITE_KEY=your_site_key
TURNSTILE_SECRET_KEY=your_secret_key

# 数据库
DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_USER=root
DATABASE_PASSWORD=rootpwd
DATABASE_NAME=sellcard

# JWT
JWT_SECRET=your_secret_key
JWT_TTL_MINUTES=1440

# 更多变量见 ENV_SETUP.md
```

## 功能特性

- ✅ 基于 JWT 的身份验证
- ✅ Turnstile 人机验证（可选）
- ✅ 响应式后台管理界面（Vuetify）
- ✅ 导航抽屉 + 主内容区域布局
- ✅ Material Design 设计规范
- ✅ 支持环境变量配置
- ✅ TypeScript 类型支持
- ✅ Swagger API 文档

## 文件说明

| 文件 | 说明 |
|-----|-----|
| `configs/config.yaml` | ❌ 本地配置（已忽略） |
| `configs/config.example.yaml` | ✅ 配置示例 |
| `.gitignore` | ✅ Git 忽略规则 |
| `ENV_SETUP.md` | 📖 环境配置详细指南 |
| `TURNSTILE_SETUP.md` | 📖 Turnstile 集成指南 |
| `start-dev.sh` | 🚀 Linux/macOS 快速启动 |
| `start-dev.bat` | 🚀 Windows 快速启动 |

## Git 提交

### ✅ 可以提交
- `.env.example` - 环境变量示例
- `config.example.yaml` - 配置示例
- `.gitignore` - 忽略规则
- 所有源代码文件

### ❌ 不能提交
- `.env.local` - 本地环境变量
- `config.yaml` - 本地配置
- 任何包含敏感信息文件

## 常见问题

### Q: 如何启用 Turnstile 人机验证？

A: 在生产环境设置以下环境变量：
```bash
export TURNSTILE_ENABLED=true
export TURNSTILE_SITE_KEY=your_site_key
export TURNSTILE_SECRET_KEY=your_secret_key
```
详见 **[TURNSTILE_SETUP.md](server/TURNSTILE_SETUP.md)**

### Q: 如何修改数据库连接？

A: 使用环境变量覆盖：
```bash
export DATABASE_HOST=your_host
export DATABASE_PASSWORD=your_password
```
或修改 `server/configs/config.yaml`

### Q: 如何停止所有服务？

A: 
- 按 `Ctrl+C` 关闭前端和后端窗口
- 停止 MySQL 容器：
  ```bash
  cd database
  docker-compose down
  ```

### Q: MySQL 连接出错怎么办？

A: 检查几点：
1. Docker 是否运行：`docker ps`
2. MySQL 容器状态：`docker-compose ps`
3. 等待 MySQL 启动（通常需要 10-15 秒）
4. 查看日志：`docker-compose logs db`

详见 **[ENV_SETUP.md](ENV_SETUP.md#常见问题)**

### Q: 项目无法启动？

A: 检查以下几点：
1. 已安装所需工具（Go、Node.js、Docker、Docker Compose）
2. MySQL 容器已启动：`docker-compose up -d`
3. 依赖已安装：`go mod tidy` 和 `npm install`
4. 相关端口未被占用（3306、8080、5173）

## 相关链接

- [Gin 文档](https://gin-gonic.com/)
- [Vue 3 文档](https://vuejs.org/)
- [Vuetify 文档](https://vuetifyjs.com/)
- [Turnstile 文档](https://developers.cloudflare.com/turnstile/)

## 许可证

MIT