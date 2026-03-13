package models

import (
	"time"
)

type User struct {
	UserId     int64   `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Email      string  `json:"email" binding:"required" gorm:"unique;not null"`
	Password   *string `json:"password" binding:"required"`
	GoogleId   *string
	Name       string    `json:"name" binding:"required" gorm:"not null"`
	Age        *int64    `json:"age"`
	Gender     string    `json:"gender"`
	ProfilePic *string   `json:"profile_pic"`
	CreatedAt  time.Time //YYYY-MM-DD HH:MM:SS.microseconds stored in DB
	UpdatedAt  time.Time
}

type UserUsageMV struct {
	UserID           int64 `gorm:"column:user_id;primaryKey;"`
	TotalStorageUsed int64 `gorm:"column:total_storage_used;"`
}
