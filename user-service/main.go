package main

import (
    "github.com/clim-bot/user-service/config"
    "github.com/clim-bot/user-service/models"
    "github.com/clim-bot/user-service/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    db := config.InitDB()
    db.AutoMigrate(&models.User{})

    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run()
}
