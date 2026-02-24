# Cloudflare Turnstile 集成指南

## 概述

本项目已集成 Cloudflare Turnstile（人机验证工具），用于增强登录页面的安全性。

## 配置说明

### 1. 获取 Turnstile 凭证

1. 访问 [Cloudflare Dashboard](https://dash.cloudflare.com/)
2. 登录或创建账户
3. 进入 **Turnstile** 部分
4. 创建新的 Turnstile 站点，获取以下凭证：
   - **Site Key**（公钥）
   - **Secret Key**（私钥）

### 2. 后端配置

编辑 `/server/configs/config.yaml`：

```yaml
turnstile:
  enabled: true                          # 启用/禁用 Turnstile
  site_key: "0x4AAAAAAAv8X5OMVd5YPvYA"  # 替换为你的 Site Key
  secret_key: "0x4AAAAAAAv8X3-WIkvFfOiGkXz0Pj0cFh"  # 替换为你的 Secret Key
```

**字段说明：**
- `enabled`: 是否启用 Turnstile 验证（`true`/`false`）
- `site_key`: Turnstile 公钥，用于前端渲染验证组件
- `secret_key`: Turnstile 私钥，仅用于后端服务器验证

### 3. 前端配置

编辑 `/web/.env.local`：

```env
VITE_API_URL=http://localhost:8080
VITE_TURNSTILE_SITE_KEY=0x4AAAAAAAv8X5OMVd5YPvYA  # 替换为你的 Site Key
```

## 工作流程

### 用户登录流程

1. **前端**：页面加载时，加载 Cloudflare Turnstile 脚本
2. **前端**：用户完成人机验证，获得验证 token
3. **前端**：用户输入用户名、密码并提交登录表单
4. **前端**：将用户名、密码、Turnstile token 发送到后端
5. **后端**：验证 Turnstile token 的有效性（向 Cloudflare API 验证）
6. **后端**：如果 Turnstile 验证通过，继续验证用户凭证
7. **后端**：验证成功后返回 JWT token

### 禁用 Turnstile 验证

如果需要禁用 Turnstile 验证（仅用于开发），修改配置：

```yaml
turnstile:
  enabled: false
```

此时，前端将不渲染验证组件，后端也将跳过验证步骤。

## API 接口

### 登录端点

**URL**: `POST /api/v1/auth/login`

**请求体**：

```json
{
  "username": "admin",
  "password": "123456",
  "turnstile_token": "token_from_cloudflare"  // 如果启用了 Turnstile，此字段必需
}
```

**响应**（成功 200）：

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "username": "admin",
    "role": "admin"
  }
}
```

**错误响应**：

- `400 Bad Request`: 请求参数无效或 Turnstile token 为空
- `401 Unauthorized`: Turnstile 验证失败或登录凭证错误
- `500 Internal Server Error`: 服务器内部错误

## 环境变量

### 后端

- `TURNSTILE_ENABLED`: 启用/禁用 Turnstile (配置文件中设置)
- `TURNSTILE_SITE_KEY`: Turnstile 公钥
- `TURNSTILE_SECRET_KEY`: Turnstile 私钥

### 前端

- `VITE_API_URL`: API 服务器地址（默认：`http://localhost:8080`）
- `VITE_TURNSTILE_SITE_KEY`: Turnstile 公钥

## 常见问题

### Q: Turnstile 验证失败怎么办？

A: 检查以下几点：
1. Site Key 和 Secret Key 配置是否正确
2. 确保 Turnstile 在 Cloudflare Dashboard 中已启用
3. 检查网络连接，确保可以访问 `https://challenges.cloudflare.com`
4. 查看浏览器控制台的错误日志

### Q: 如何在开发环境中禁用 Turnstile？

A: 在 `config.yaml` 中设置 `enabled: false`

### Q: Turnstile 支持哪些验证模式？

A: 当前实现支持互动模式（Managed Challenge），但可以根据需要扩展支持其他模式。

## 安全提示

⚠️ **重要**：
- 不要在公开的代码库中泄露 Secret Key
- Secret Key 仅应在后端服务器上使用
- 定期轮换 Secret Key
- 在生产环境中启用 Turnstile 验证

## 参考资源

- [Cloudflare Turnstile 文档](https://developers.cloudflare.com/turnstile/)
- [Turnstile API 参考](https://developers.cloudflare.com/turnstile/get-started/)
