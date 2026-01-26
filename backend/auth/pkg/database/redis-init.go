package database

import (
	"context"
	"github.com/gofiber/storage/redis"
	"time"
)

type RedisClient struct {
	*redis.Storage
}

func RedisInit() *RedisClient {
	db := redis.New(redis.Config{
		Host: "redis",
		Port: 6379,
	})

	return &RedisClient{
		Storage: db,
	}
}

func (r *RedisClient) GetWithContext(ctx context.Context, key string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return r.Get(key)
}

func (r *RedisClient) SetWithContext(ctx context.Context, key string, val []byte, exp time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return r.Set(key, val, exp)
}

func (r *RedisClient) DeleteWithContext(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return r.Delete(key)
}

func (r *RedisClient) ResetWithContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return r.Reset()
}
