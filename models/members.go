package models

import (
	"gorm.io/gorm"
)

type Members struct {
	gorm.Model
	Name         string `json:"Name"`
	EmailAddress string `json:"EmailAddress"`
	Password     string `json:"Password"`
}
