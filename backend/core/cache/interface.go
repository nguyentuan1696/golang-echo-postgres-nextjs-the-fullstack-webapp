package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICache interface {
	GetClient() *redis.Client
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Close() error
	IsLoginBlocked(ctx context.Context, key string) (bool, error)
	IncrementLoginAttempt(ctx context.Context, key string) error
}
