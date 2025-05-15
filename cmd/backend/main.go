package main

import (
    "log"
    "backend/internal/config"
    "backend/internal/db"
    "backend/internal/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    cfg, err := config.LoadConfig("config/config.yaml")
    if err != nil {
        log.Fatalf("載入設定失敗: %v", err)
    }

    mysqlDB, err := db.InitMySQL(
        cfg.Database.Username,
        cfg.Database.Password,
        cfg.Database.Host,
        cfg.Database.Name,
    )
    if err != nil {
        log.Fatalf("資料庫連接失敗: %v", err)
    }

    r := gin.Default()

    r.GET("/tables", handler.GetTables(mysqlDB))

    r.Run(":8080")
}

