package models

import (
	"time"
)

type User struct {
	UserId     int64   `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Email      string  `json:"email" binding:"required" gorm:"unique;not null"`
	Password   *string `json:"password" binding:"required"`
	GoogleId   *string
	Name       string  `json:"name" binding:"required" gorm:"not null"`
	Age        *int64  `json:"age"`
	Gender     string  `json:"gender"`
	ProfilePic *string `json:"profile_pic"`

	IsVerified   bool    `json:"is_verified" gorm:"default:false"`
	OTP          *string `json:"-"` // hide from API
	OTPExpiresAt *time.Time

	CreatedAt time.Time //YYYY-MM-DD HH:MM:SS.microseconds stored in DB
	UpdatedAt time.Time
}
