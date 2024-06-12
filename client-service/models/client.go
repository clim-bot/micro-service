package models

import (
    "gorm.io/gorm"
)

type Client struct {
    gorm.Model
    Name      string `json:"name" binding:"required"`
    Email     string `json:"email" binding:"required"`
    Phone     string `json:"phone" binding:"required"`
    ZipCode   string `json:"zipcode" binding:"required"`
}
