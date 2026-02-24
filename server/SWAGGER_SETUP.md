# Swagger API Documentation Setup

## 概览

本项目已集成 Swagger (OpenAPI) 文档，支持通过配置文件启用/禁用 Swagger UI。

## 配置

### 启用/禁用 Swagger

编辑 `configs/config.yaml` 文件中的 Swagger 配置：

```yaml
swagger:
  enabled: true   # 设置为 true 启用，false 禁用
```

## 使用方法

### 1. 启动服务器

```bash
cd /workspaces/sellcard/server
go run cmd/server/main.go
```

### 2. 访问 Swagger UI

当 Swagger 启用时，访问以下 URL：

```
http://localhost:8080/swagger/index.html
```

## 生成 Swagger 文档

如果修改了 API 的处理函数注释，需要重新生成  Swagger 文档：

```bash
cd /workspaces/sellcard/server
go run github.com/swaggo/swag/cmd/swag init -g cmd/server/main.go
```

## Swagger 注释规范

### API 端点注释示例

```go
// @Summary        用户登录
// @Description    通过用户名和密码登录，获取 JWT token
// @Tags           auth
// @Accept         json
// @Produce        json
// @Param          request body LoginRequest true "登录凭证"
// @Success        200 {object} map[string]interface{} "登录成功，返回 token"
// @Failure        400 {object} map[string]interface{} "请求无效"
// @Failure        401 {object} map[string]interface{} "凭证无效"
// @Failure        500 {object} map[string]interface{} "服务器错误"
// @Router         /auth/login [post]
func Login(cfg *config.Config) gin.HandlerFunc {
    // ...
}
```

## 项目结构

```
server/
├── cmd/server/
│   └── main.go              # 应用入口（包含 Swagger 全局注释）
├── configs/
│   └── config.yaml          # 配置文件（包含 swagger 部分）
├── docs/                    # Swagger 自动生成的文档目录
│   ├── docs.go             # Go 代码形式的 Swagger 定义
│   ├── swagger.json        # JSON 格式的 Swagger 定义
│   └── swagger.yaml        # YAML 格式的 Swagger 定义
├── internal/
│   ├── config/
│   │   └── config.go       # 配置结构（包含 Swagger 配置）
│   ├── controller/
│   │   └── auth_controller.go  # 控制器（包含 Swagger 注释）
│   └── routes/
│       └── routes.go       # 路由定义（集成 Swagger 路由）
└── go.mod
```

## 主要依赖

- `github.com/swaggo/swag` - Swagger 生成器
- `github.com/swaggo/gin-swagger` - Gin 框架的 Swagger 处理器
- `github.com/swaggo/files` - Swagger 静态文件处理

## 注意事项

1. **安全性**：生产环境建议将 `swagger.enabled` 设置为 `false`
2. **重新生成文档**：修改 API 注释后，需要运行 `swag init` 命令
3. **API 基础路径**：当前设置为 `/api/v1`，Swagger UI 将显示相对于此路径的端点
4. **认证**：API 已配置 Bearer Token 认证，在 Swagger UI 中可使用"Authorize"按钮添加 token

## 故障排除

### Swagger UI 无法加载

确保：
1. `swagger.enabled` 设置为 `true`
2. `docs/docs.go` 文件存在且内容正确
3. 运行 `swag init` 重新生成文档

### 某些端点在 Swagger UI 中不显示

确保：
1. 端点函数有正确的 Swagger 注释
2. 运行 `swag init` 重新生成文档
3. 重启服务器
