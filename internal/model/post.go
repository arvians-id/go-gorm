package model

import (
	"time"
)

type Post struct {
	ID        uint64     `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"type:varchar(255);not null" json:"title"`
	Body      string     `gorm:"type:text;not null" json:"body"`
	Comments  []*Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type PostResponse struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	CommentBody string    `json:"comment_body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
