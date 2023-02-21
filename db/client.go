package db

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type Client struct {
	Ctx    context.Context
	Client *redis.Client
}
