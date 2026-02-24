package routes

import (
    "sellcard/server/internal/config"
    "sellcard/server/internal/controller"
    "sellcard/server/internal/middleware"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "sellcard/server/docs"
)

func RegisterRoutes(r *gin.Engine) {
    cfg, _ := config.LoadConfig("./configs/config.yaml")

    // 根据配置启用 Swagger 文档
    if cfg.Swagger.Enabled {
        r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }

    // API v1 路由组
    api := r.Group("/api/v1")
    {
        // 认证相关路由
        api.POST("/auth/login", controller.Login(cfg))
    }

    // 管理员路由（需要 JWT 认证）
    admin := api.Group("/admin")
    admin.Use(middleware.JWTAuth(cfg))
    {
        // 管理员仪表盘
        admin.GET("/dashboard", func(c *gin.Context) {
            c.JSON(200, gin.H{"msg": "管理员仪表盘"})
        })
    }
}
