package models

import (
    "gorm.io/gorm"
)

type Order struct {
    gorm.Model
    ClientID uint   `json:"client_id" binding:"required"`
    Product  string `json:"product" binding:"required"`
    Amount   int    `json:"amount" binding:"required"`
}
