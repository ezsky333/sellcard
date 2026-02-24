package main

import (
    "log"
    "sellcard/server/internal/config"
    "sellcard/server/internal/model"
    "sellcard/server/internal/repository"
    "sellcard/server/internal/routes"
    "time"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// @title 售卡系统 API
// @version 1.0
// @description 售卡系统身份认证 API 服务器
// @termsOfService http://swagger.io/terms/

// @contact.name API 支持
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
    cfg, err := config.LoadConfig("./configs/config.yaml")
    if err != nil {
        log.Fatalf("load config: %v", err)
    }
    dsn := cfg.DatabaseDSN()
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect db: %v", err)
    }

    // create table if not exists (avoid destructive alter operations by AutoMigrate)
    if !db.Migrator().HasTable(&model.User{}) {
        if err := db.Migrator().CreateTable(&model.User{}); err != nil {
            log.Fatalf("create table: %v", err)
        }
    }

    // set up repository instances
    repository.SetDB(db)

    r := gin.Default()
    r.Use(func(c *gin.Context) {
        // simple request start time header
        c.Writer.Header().Set("X-Server-Time", time.Now().Format(time.RFC3339))
        c.Next()
    })

    routes.RegisterRoutes(r)

    addr := cfg.ServerAddress()
    if err := r.Run(addr); err != nil {
        log.Fatalf("server run: %v", err)
    }
}
