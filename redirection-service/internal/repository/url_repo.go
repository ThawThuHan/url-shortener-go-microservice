package repository

import (
	"context"
	"encoding/json"
	"log"
	"redirection-service/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type URLRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewURLRepository(db *gorm.DB, redis *redis.Client) *URLRepository {
	return &URLRepository{db: db, redis: redis}
}

func (r *URLRepository) GetOriginURL(ctx context.Context, shortCode string) (*model.URL, error) {
	var url model.URL
	result, err := r.redis.Get(ctx, "url:"+shortCode).Result()
	if err == redis.Nil || err != nil {
		if err := r.db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
			return nil, err
		}
		jsonString, err := json.Marshal(url)
		if err != nil {
			log.Println("json marshaling error")
		}
		if err := r.redis.Set(ctx, "url:"+shortCode, jsonString, 1*time.Hour).Err(); err != nil {
			log.Println("redis cache set error")
		}
	}

	if result != "" {
		json.Unmarshal([]byte(result), &url)
	}

	return &url, nil
}
