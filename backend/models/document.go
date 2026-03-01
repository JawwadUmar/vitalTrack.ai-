package models

import "time"

type Document struct {
	ID         string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     int64     `json:"user_id"`
	FileID     string    `json:"file_id"`
	Category   string    `json:"category"`
	ReportType string    `json:"report_type"`
	FileType   string    `json:"file_type"`
	Tags       string    `json:"tags"` // JSON string
	Status     string    `json:"status"`
	ReportDate time.Time `json:"report_date"`
}

type CalendarRequest struct {
	Month    int      `json:"month"`
	Year     int      `json:"year"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}
