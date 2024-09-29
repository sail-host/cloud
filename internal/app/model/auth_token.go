package model

import "time"

type AuthToken struct {
	BaseModel
	Token     string    `json:"token" gorm:"varchar(255);not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	ExpiredAt time.Time `json:"expired_at" gorm:"not null"`
}
