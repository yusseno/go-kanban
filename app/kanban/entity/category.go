package entity

import "gorm.io/gorm"

type Category struct {
	*gorm.DB
	role_id   string
	name      string
	deskripsi string
}
