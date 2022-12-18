package model

type Role struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Role  string  `gorm:"type:varchar(255);not null;unique" json:"role"`
	Users []*User `gorm:"many2many:user_roles" json:"user_roles"`
}
