package database

import (
	"context"
	"time"
)

type LimiterType int

const (
	TypeIP LimiterType = iota
	TypeKey
)

// RateInterface contains all methods used by middleware to check stored data.
// This interface can be implemented by any type of database, it will not affect the middleware behavior.
type RateInterface interface {
	NewRequest(ctx context.Context, limiterType LimiterType, val string, ttl time.Duration) error
	FindRequests(ctx context.Context, limiterType LimiterType, val string) (int, error)

	Block(ctx context.Context, limiterType LimiterType, val string, ttl time.Duration) error
	FindBlocker(ctx context.Context, limiterType LimiterType, val string) (bool, error)
}
