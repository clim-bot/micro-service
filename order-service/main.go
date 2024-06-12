package main

import (
    "github.com/clim-bot/order-service/config"
    "github.com/clim-bot/order-service/models"
    "github.com/clim-bot/order-service/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    db := config.InitDB()
    db.AutoMigrate(&models.Order{})

    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run("8081")
}
