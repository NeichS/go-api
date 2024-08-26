package models

import (
	"gorm.io/gorm"	
)

type User struct {
	gorm.Model

	nombre string `gorm:"not null"`
	apellido string `gorm:"not null"`
	email string `gorm:"not null;unique_index"`
	Tareas []Task
}