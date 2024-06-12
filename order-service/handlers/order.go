package handlers

import (
    "net/http"
    "github.com/clim-bot/order-service/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB
}

func (h *Handler) CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (h *Handler) GetOrders(c *gin.Context) {
    var orders []models.Order
    if err := h.DB.Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}

func (h *Handler) GetOrder(c *gin.Context) {
    orderID := c.Param("id")
    var order models.Order
    if err := h.DB.First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (h *Handler) DeleteOrder(c *gin.Context) {
    orderID := c.Param("id")
    if err := h.DB.Delete(&models.Order{}, orderID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
