package cache

import (
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Client *redis.Client
}

func (c *Client) InitializeConnection(host string, passWord string, poolSize, minIdleConns, DB int) {
	c.Client = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     passWord,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
		DB:           0,
	})
}

// Pipeline .
func (c *Client) Pipeline() redis.Pipeliner {
	return c.Client.Pipeline()
}
