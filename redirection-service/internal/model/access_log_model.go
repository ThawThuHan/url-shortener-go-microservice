package model

import "time"

type AccessLog struct {
	ID         uint      `json:"id" gorm:"primary_key;autoIncrement;not null"`
	ShortURLID uint      `json:"short_url_id" gorm:"not null"`
	IPAddress  string    `json:"ip_address" gorm:"not null"`
	Location   string    `json:"location" gorm:"not null"`
	City       string    `json:"city" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
