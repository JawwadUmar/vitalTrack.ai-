package models

import "time"

type OneTimePassword struct {
	Id           int64 `json:"user_id" gorm:"column:user_id;primaryKey"`
	Email        string
	OTP          *string `json:"-"` // hide from API
	OTPExpiresAt *time.Time
	User         User `gorm:"foreignKey:Id;references:UserId;constraint:OnDelete:CASCADE"`
}
