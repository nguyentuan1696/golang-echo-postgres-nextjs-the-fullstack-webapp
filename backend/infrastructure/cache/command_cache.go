package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"strings"
	"thichlab-backend-docs/constant"
	"time"
)

// DelKeys example: https://github.com/redis/go-redis/tree/master/example/del-keys-without-ttl
func (c *Client) DelKeys(ctx context.Context, keyCache string) error {
	pattern := keyCache + "*"
	iter := c.Client.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		if err := c.Client.Del(ctx, key).Err(); err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}

// HSet Sets the specified fields to their respective values in the hash stored at key.
func (c *Client) HSet(ctx context.Context, key string, field string, value any) error {
	return c.Client.HSet(ctx, key, field, value).Err()
}

// HGet Returns the value associated with field in the hash stored at key.
func (c *Client) HGet(ctx context.Context, key, field string) error {
	return c.Client.HGet(ctx, key, field).Err()
}

// Exists Returns if key exists.
func (c *Client) Exists(ctx context.Context, key string) int64 {
	return c.Client.Exists(ctx, key).Val()
}

// Set key to hold the string value. If key already holds a value, it is overwritten, regardless of its type. Any previous time to live associated with the key is discarded on successful SET operation.
func (c *Client) Set(ctx context.Context, key string, value any, expiration time.Duration) string {
	out, err := json.Marshal(value)
	if err != nil && !strings.Contains(err.Error(), "redis: nil") {
		return ""
	}

	err = c.Client.Set(ctx, key, out, expiration).Err()
	if err != nil {
		return ""
	}
	return c.Client.Set(ctx, key, value, expiration).Val()
}

// Get the value of key. If the key does not exist the special value nil is returned. An error is returned if the value stored at key is not a string, because GET only handles string values.
func (c *Client) Get(ctx context.Context, key string) string {
	return c.Client.Get(ctx, key).Val()
}

func (c *Client) Del(ctx context.Context, key string) int64 {
	return c.Client.Del(ctx, key).Val()
}

func (c *Client) ZIncrBy(ctx context.Context, key string, increment float64, member string) error {
	return c.Client.ZIncrBy(ctx, key, increment, member).Err()
}
func (c *Client) GetHighestScore(ctx context.Context, key string, pageSize, pageIndex int64) []string {
	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = 5
		pageIndex = 0
	}
	return c.Client.ZRevRangeByScore(ctx, key, &redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: pageIndex*pageSize - pageSize, Count: pageSize}).Val()
}
