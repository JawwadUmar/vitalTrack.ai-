package models

import "time"

type File struct {
	ID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OriginalName string
	StoredName   string
	S3Key        string
	FileSize     int64
	MimeType     string
	UploadedBy   string
	CreatedAt    time.Time
}
