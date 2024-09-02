package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang-template-service/config"
)

type RedisClient struct {
	RedisClient *redis.Client
}

// New constructs new DatabaseConnection
func New(config config.RedisConfig) (*RedisClient, error) {
	connStr := Connect(config)

	return &RedisClient{
		RedisClient: connStr,
	}, nil
}

func Connect(config config.RedisConfig) *redis.Client {
	var (
		rdb *redis.Client
	)

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedistHost, config.RedisPort),
		Password: config.RedisPass, // no password set
		DB:       config.RedisDb,   // use default DB
	})

	return rdb
}
