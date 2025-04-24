package model

import "time"

type URL struct {
	ID        uint      `json:"id" gorm:"primary_key;autoIncrement;not null"`
	SessionID string    `json:"session_id" gorm:"not null"`
	OriginURL string    `json:"origin_url" gorm:"unique;not null"`
	ShortCode string    `json:"short_code" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
