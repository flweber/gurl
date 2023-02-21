package db

import (
	"context"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func NewClient(addr, pass string, db int) *Client {
	client := Client{}
	client.Ctx = context.Background()
	client.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	return &client
}
