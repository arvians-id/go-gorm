package model

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"unique;type:varchar(255);not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Roles     []*Role   `gorm:"many2many:user_roles" json:"user_roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
