package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"type:varchar(200);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(200);not null" json:"last_name"`
	Email     string `gorm:"type:varchar(200);not null;unique_index" json:"email"`
	Tasks     []Task ` json:"tasks"`
}
