package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	IpStoragePrefix  = "ip-"
	KeyStoragePrefix = "key-"
	IpBlockPrefix    = "ip-blocked-"
	KeyBlockPrefix   = "key-blocked-"
)

type RateLimitStorage struct {
	Client *redis.Client
}

func NewRateLimitStorage(client *redis.Client) RateInterface {
	return &RateLimitStorage{Client: client}
}

func (r *RateLimitStorage) NewRequest(ctx context.Context, limiterType LimiterType, val string, ttl time.Duration) error {
	redisKey := getRequestKey(limiterType, val)
	requests, err := r.FindRequests(ctx, limiterType, val)
	if err != nil {
		return err
	}

	var value int
	if requests != 0 {
		value = requests
	}
	value++

	return r.Client.Set(ctx, redisKey, value, ttl).Err()
}

func (r *RateLimitStorage) FindRequests(ctx context.Context, limiterType LimiterType, val string) (int, error) {
	redisKey := getRequestKey(limiterType, val)
	requests, err := r.Client.Get(ctx, redisKey).Int()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return 0, err
		}
	}
	return requests, nil
}

func getRequestKey(limiterType LimiterType, val string) string {
	if limiterType == TypeIP {
		return fmt.Sprintf("%s%s", IpStoragePrefix, val)
	}
	return fmt.Sprintf("%s%s", KeyStoragePrefix, val)
}

func (r *RateLimitStorage) Block(ctx context.Context, limiterType LimiterType, val string, ttl time.Duration) error {
	redisKey := getBlockKey(limiterType, val)
	block, err := r.FindBlocker(ctx, limiterType, val)
	if err != nil {
		return err
	}

	if block {
		return nil
	}

	return r.Client.Set(ctx, redisKey, true, ttl).Err()
}

func (r *RateLimitStorage) FindBlocker(ctx context.Context, limiterType LimiterType, val string) (bool, error) {
	redisKey := getBlockKey(limiterType, val)
	found, err := r.Client.Get(ctx, redisKey).Bool()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return false, err
		}
	}
	return found, nil
}

func getBlockKey(limiterType LimiterType, val string) string {
	if limiterType == TypeIP {
		return fmt.Sprintf("%s%s", IpBlockPrefix, val)
	}
	return fmt.Sprintf("%s%s", KeyBlockPrefix, val)
}
