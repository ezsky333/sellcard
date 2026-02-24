package controller

import (
	"net/http"
	"sellcard/server/internal/config"
	"sellcard/server/internal/service"
	"sellcard/server/internal/utils"

	"github.com/gin-gonic/gin"
)

// 登录请求参数
type LoginRequest struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	TurnstileToken string `json:"turnstile_token"`
}

// @Summary 用户登录
// @Description 使用用户名和密码登录，获取 JWT token
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body LoginRequest true "登录凭证"
// @Success 200 {object} map[string]interface{} "登录成功，返回 token 和用户信息"
// @Failure 400 {object} map[string]interface{} "请求参数无效"
// @Failure 401 {object} map[string]interface{} "登录凭证错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /auth/login [post]
func Login(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效"})
			return
		}

		// 如果启用了 Turnstile 验证，则验证 token
		if cfg.Turnstile.Enabled {
			if req.TurnstileToken == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Turnstile token 不能为空"})
				return
			}

			verified, err := utils.VerifyTurnstile(cfg.Turnstile.SecretKey, req.TurnstileToken)
			if err != nil || !verified {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Turnstile 验证失败"})
				return
			}
		}

		// 验证用户身份
		u, err := service.Authenticate(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录凭证错误"})
			return
		}
		// 生成 JWT token
		token, err := utils.GenerateToken(cfg.JWT.Secret, u.Username, u.Role, cfg.JWT.TTLMinutes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token 生成失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "user": gin.H{"username": u.Username, "role": u.Role}})
	}
}
