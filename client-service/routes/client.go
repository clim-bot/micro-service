package routes

import (
    "github.com/clim-bot/client-service/handlers"
    "github.com/clim-bot/client-service/utils"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        c.Set("userID", uint(claims["userID"].(float64)))
        c.Next()
    }
}

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    h := &handlers.Handler{DB: db}

    r.POST("/clients", authMiddleware(), h.CreateClient)
    r.PUT("/clients/:id", authMiddleware(), h.UpdateClient)
    r.DELETE("/clients/:id", authMiddleware(), h.DeleteClient)
    r.GET("/clients/:id", authMiddleware(), h.GetClient)
    r.GET("/clients", authMiddleware(), h.GetClients)
}
