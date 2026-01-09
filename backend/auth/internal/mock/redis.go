package mock

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"time"
)

type Redis struct {
	fiber.Storage
	store map[interface{}]interface{}
}

func NewMockRedis() *Redis {
	return &Redis{
		store: make(map[interface{}]interface{}),
	}
}

func (r *Redis) GetWithContext(ctx context.Context, key string) ([]byte, error) {
	return r.Get(key)
}

func (r *Redis) SetWithContext(ctx context.Context, key string, val []byte, exp time.Duration) error {
	return r.Set(key, val, exp)
}

func (r *Redis) DeleteWithContext(ctx context.Context, key string) error {
	return nil
}

func (r *Redis) ResetWithContext(ctx context.Context) error {
	return nil
}

func (r *Redis) Get(key string) ([]byte, error) {
	return r.store[key].([]byte), nil
}

func (r *Redis) Set(key string, val []byte, exp time.Duration) error {
	r.store[key] = val
	return nil
}

func (r *Redis) Delete(key string) error {
	return nil
}

func (r *Redis) Reset() error {
	return nil
}

func (r *Redis) Close() error {
	return nil
}
