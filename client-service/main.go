package main

import (
	"github.com/clim-bot/client-service/config"
    "github.com/clim-bot/client-service/models"
    "github.com/clim-bot/client-service/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    db := config.InitDB()
    db.AutoMigrate(&models.Client{})

    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run("8082")
}
