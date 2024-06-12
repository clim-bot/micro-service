package routes

import (
    "github.com/clim-bot/user-service/handlers"
    "github.com/clim-bot/user-service/utils"
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

    r.POST("/register", h.Register)
    r.POST("/login", h.Login)
    r.POST("/logout", authMiddleware(), h.Logout)
    r.POST("/update-password", authMiddleware(), h.UpdatePassword)
    r.POST("/reset-password", h.ResetPassword)
    r.GET("/profile", authMiddleware(), h.Profile)
}
