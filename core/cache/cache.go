package cache

import (
	"context"
	"go-api-starter/core/logger"
	"sync"
	"time"

	"go-api-starter/core/constants"

	"github.com/redis/go-redis/v9"
)

var (
	instance *Cache
	once     sync.Once
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr, password string, db int) *Cache {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		// Test the connection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := client.Ping(ctx).Result()
		if err != nil {
			logger.Error("Failed to connect to Redis: " + err.Error())
		}

		instance = &Cache{
			client: client,
		}
	})
	return instance
}

// GetClient returns the Redis client
func (c *Cache) GetClient() *redis.Client {
	return c.client
}

// Set sets a key-value pair with expiration
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func (c *Cache) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.client.Get(ctx, key)
}

// Del removes a key
func (c *Cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Incr increments a key's value
func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

// Expire sets expiration for a key
func (c *Cache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

// Close closes the Redis connection
func (c *Cache) Close() error {
	return c.client.Close()
}

func (c *Cache) IsLoginBlocked(ctx context.Context, key string) (bool, error) {
	count, err := c.client.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return false, err
	}
	if count >= constants.MaxLoginAttempts {

		return true, nil
	}
	return false, nil
}

func (c *Cache) IncrementLoginAttempt(ctx context.Context, key string) error {
	// Tăng +1, nếu chưa tồn tại thì set TTL luôn
	val, err := c.client.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	if val == 1 {
		// Set TTL nếu là lần đầu tiên
		err = c.client.Expire(ctx, key, constants.BlockDuration).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
