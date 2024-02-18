package models

import (
	"gorm.io/gorm"
)

type Members struct {
	gorm.Model
	Identity     int    `json:"Identity"`
	Name         string `json:"Name"`
	Email_addres string `json:"Email_addres"`
	Password     string `json:"Password"`
}
