package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Fullname string
	Email    string
	Password string
	Is_Role  int
}

type ResLogin struct {
	Email   string
	Message string
}
