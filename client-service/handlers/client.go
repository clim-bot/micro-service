package handlers

import (
    "net/http"
    "github.com/clim-bot/client-service/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB
}

func (h *Handler) CreateClient(c *gin.Context) {
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&client).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client"})
        return
    }

    c.JSON(http.StatusOK, client)
}

func (h *Handler) UpdateClient(c *gin.Context) {
    id := c.Param("id")
    var client models.Client
    if err := h.DB.First(&client, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
        return
    }

    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Save(&client).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client"})
        return
    }

    c.JSON(http.StatusOK, client)
}

func (h *Handler) DeleteClient(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Client{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}

func (h *Handler) GetClient(c *gin.Context) {
    id := c.Param("id")
    var client models.Client
    if err := h.DB.First(&client, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
        return
    }

    c.JSON(http.StatusOK, client)
}

func (h *Handler) GetClients(c *gin.Context) {
    var clients []models.Client
    if err := h.DB.Find(&clients).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
        return
    }

    c.JSON(http.StatusOK, clients)
}
