package database

import (
	"context"
	"redirection-service/configs"
	"time"

	"github.com/redis/go-redis/v9"
)

func Redis(cfg configs.Config) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis_Host + ":" + cfg.Redis_Port,
		Password: cfg.Redis_Password,
		DB:       0,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
