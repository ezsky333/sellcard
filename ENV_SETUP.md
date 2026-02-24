# 环境配置指南

本项目采用**环境变量注入**方案管理敏感配置，支持本地开发和生产部署的无缝切换。

## 目录结构

```
.
├── .gitignore              # 项目根级忽略规则
├── server/
│   ├── .gitignore         # 后端忽略规则（忽略 config.yaml）
│   ├── configs/
│   │   ├── config.yaml           # 开发配置（禁用 Turnstile）
│   │   └── config.example.yaml   # 配置示例和说明
│   └── internal/config/config.go # 支持环境变量覆盖
└── web/
    ├── .gitignore       # 前端忽略规则（忽略 .env.local）
    ├── .env.local       # 本地环境变量（不提交）
    └── .env.example     # 环境变量示例（已提交）
```

## 本地开发（推荐方式）

### 数据库启动

1. **启动 MySQL（Docker）**：

```bash
cd database
docker-compose up -d
```

MySQL 服务将在 `localhost:3306` 启动，凭证：
- 用户: `root`
- 密码: `rootpwd`
- 数据库: `sellcard`

初始化 SQL 脚本会自动执行：`/server/scripts/db/init.sql`

### 后端开发

1. **默认配置已禁用 Turnstile**，直接运行：

```bash
cd server
go run cmd/server/main.go
```

服务器启动在 `http://localhost:8080`

> 注意：确保 MySQL 已启动，否则后端会失败

### 前端开发

1. **.env.local 已配置好**，直接运行：

```bash
cd web
npm install
npm run dev
```

开发服务器启动在 `http://localhost:5173`

> 注意：`.env.local` 已在 `.gitignore` 中，不会被提交到 GitHub

### 快速启动（自动化）

或者使用一键启动脚本（自动启动 MySQL、后端、前端）：

```bash
# Linux/macOS
./start-dev.sh

# Windows
start-dev.bat
```

## 生产部署

### 方式 1：环境变量注入（推荐 - 容器部署）

```bash
# 后端
export TURNSTILE_ENABLED=true
export TURNSTILE_SITE_KEY=your_site_key
export TURNSTILE_SECRET_KEY=your_secret_key
export JWT_SECRET=your_jwt_secret
export DATABASE_PASSWORD=your_db_password
export DATABASE_HOST=your_db_host

go run cmd/server/main.go
```

### 方式 2：Docker 部署

```dockerfile
FROM golang:1.21 as builder
WORKDIR /build
COPY . .
RUN go build -o server cmd/server/main.go

FROM ubuntu:latest
COPY --from=builder /build/server /usr/local/bin/
COPY --from=builder /build/configs/config.yaml /etc/server/

EXPOSE 8080
CMD ["server"]
```

运行容器时注入环境变量：

```bash
docker run -e TURNSTILE_ENABLED=true \
           -e TURNSTILE_SITE_KEY=xxx \
           -e TURNSTILE_SECRET_KEY=xxx \
           -e JWT_SECRET=your_secret \
           -p 8080:8080 \
           your-image:latest
```

### 方式 3：修改 config.yaml（仅限非敏感部署）

如果无法使用环境变量，可直接修改 `config.yaml`：

```yaml
turnstile:
  enabled: true
  site_key: "your_site_key"
  secret_key: "your_secret_key"
```

**警告**：不要将包含敏感信息的 `config.yaml` 提交到 GitHub!

## 支持的环境变量

### 服务器

| 环境变量 | 配置项 | 默认值 |
|--------|--------|--------|
| `SERVER_HOST` | `server.host` | 0.0.0.0 |
| `SERVER_PORT` | `server.port` | 8080 |

### 数据库

| 环境变量 | 配置项 | 默认值 |
|--------|--------|--------|
| `DATABASE_HOST` | `database.host` | 127.0.0.1 |
| `DATABASE_PORT` | `database.port` | 3306 |
| `DATABASE_USER` | `database.user` | root |
| `DATABASE_PASSWORD` | `database.password` | rootpwd |
| `DATABASE_NAME` | `database.dbname` | sellcard |

### JWT

| 环境变量 | 配置项 | 默认值 |
|--------|--------|--------|
| `JWT_SECRET` | `jwt.secret` | change_this_secret |
| `JWT_TTL_MINUTES` | `jwt.ttl_minutes` | 1440 |

### Swagger

| 环境变量 | 配置项 | 默认值 |
|--------|--------|--------|
| `SWAGGER_ENABLED` | `swagger.enabled` | true |

### Turnstile

| 环境变量 | 配置项 | 默认值 |
|--------|--------|--------|
| `TURNSTILE_ENABLED` | `turnstile.enabled` | false |
| `TURNSTILE_SITE_KEY` | `turnstile.site_key` | (空) |
| `TURNSTILE_SECRET_KEY` | `turnstile.secret_key` | (空) |

## 提交到 GitHub

### 可以提交的文件

✅ `.env.example` - 环境变量示例
✅ `config.example.yaml` - 配置示例
✅ `.gitignore` - 忽略规则

### 不能提交的文件

❌ `.env.local` - 本地环境变量（自动忽略）
❌ `config.yaml` - 本地配置（自动忽略）
❌ 任何包含敏感信息的文件

## 验证配置

### 检查 git 状态

```bash
# 确保敏感文件被忽略
git status

# 不应该看到下列文件：
# - server/config.yaml
# - web/.env.local
```

### 最少的环境变量配置文件内容

如果需要分享某个环境的配置示例（不含敏感信息）：

```bash
# .env.production.example
VITE_API_URL=https://api.example.com
VITE_TURNSTILE_SITE_KEY=your_site_key_here

# 生产环境 server 配置
# DATABASE_HOST=db.example.com
# DATABASE_USER=dbuser
# JWT_SECRET=<generate_a_strong_secret>
```

## 常见问题

### Q：如何快速初始化本地开发环境？

A：推荐使用一键启动脚本：
- Linux/macOS: `./start-dev.sh`
- Windows: `start-dev.bat`

或手动启动：
1. MySQL: `cd database && docker-compose up -d`
2. 后端: `cd server && go run cmd/server/main.go`
3. 前端: `cd web && npm install && npm run dev`

### Q：没有安装 Docker 怎么办？

A：虽然推荐使用 Docker，但你也可以本地安装 MySQL：

**MacOS（Homebrew）：**
```bash
brew install mysql
brew services start mysql
mysql -uroot -e "CREATE DATABASE IF NOT EXISTS sellcard;"
mysql -uroot sellcard < server/scripts/db/init.sql
```

**Linux（Ubuntu）：**
```bash
sudo apt update
sudo apt install mysql-server
sudo mysql -uroot < server/scripts/db/init.sql
```

**Windows：**
1. 下载 [MySQL Installer](https://dev.mysql.com/downloads/mysql/)
2. 安装并配置为服务
3. 运行初始化脚本：`mysql -uroot -p < server/scripts/db/init.sql`

然后修改 `server/configs/config.yaml` 的数据库配置。

### Q：MySQL 启动失败怎么办？

A：检查以下几点：
1. Docker 是否运行：`docker ps`
2. 端口是否被占用：`lsof -i :3306`（macOS/Linux）或 `netstat -ano | findstr 3306`（Windows）
3. 查看日志：`docker-compose logs db`
4. 重启：`docker-compose down && docker-compose up -d`

### Q：连接 MySQL 时 "Connection refused" 怎么办？

A：MySQL 可能还在启动中，等待约 10-15 秒。如果仍然失败：
```bash
# 检查 MySQL 容器状态
docker ps | grep mysql

# 重启容器
docker-compose restart db

# 检查日志
docker-compose logs db
```

### Q：如何停止 MySQL？

A：
```bash
cd database
docker-compose down
```

或仅停止容器（保留数据）：
```bash
docker-compose stop
```

### Q：如何为团队成员提供配置模板？

A：
1. 提交 `.env.example`（不含敏感信息）
2. 提交 `config.example.yaml`（不含敏感信息）
3. 团队成员复制为 `.env.local` 和 `config.yaml`，填入本地或真实的值

### Q：生产部署时容易遗漏环境变量怎么办？

A：使用 CI/CD 平台（GitHub Actions、GitLab CI 等）的 Secret 管理功能：

```yaml
# GitHub Actions 示例
- name: Deploy
  run: |
    export TURNSTILE_ENABLED=${{ secrets.TURNSTILE_ENABLED }}
    export TURNSTILE_SITE_KEY=${{ secrets.TURNSTILE_SITE_KEY }}
    export TURNSTILE_SECRET_KEY=${{ secrets.TURNSTILE_SECRET_KEY }}
    # ... 启动服务
```

## 总结

- ✅ 本地开发：零配置，开箱即用
- ✅ GitHub 提交：零敏感信息泄露
- ✅ 生产部署：灵活使用环境变量注入
- ✅ 团队协作：清晰的示例和文档
