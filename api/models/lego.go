package models

import ( 
	"gorm.io/gorm" 
	) 

type Lego struct {
	gorm.Model
	Lego_name   string `json:"lego_name"`
	Description string `json:"description"`
	Member_id   int    `json:"member_id"`
}
