package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// ICache interface for plain cache
type ICache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	SetObject(ctx context.Context, key string, val interface{}, duration time.Duration) error
	Delete(ctx context.Context, key string) error
	SetWithDuration(ctx context.Context, key string, value []byte, duration time.Duration) error
	SetExpireTime(ctx context.Context, key string, seconds int64) error
	Exists(ctx context.Context, keys ...string) int64

	LSet(ctx context.Context, key string, vals []byte) error
	RSet(ctx context.Context, key string, val []byte) error
	LLen(ctx context.Context, key string) (int64, error)
	LGet(ctx context.Context, key string) ([]byte, error)
	LList(ctx context.Context, key string) ([]string, error)
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	DecrBy(ctx context.Context, key string, value int64) (int64, error)
	IncrBy(ctx context.Context, key string, value int64) (int64, error)

	SetVal(ctx context.Context, key string, value string) error
	GetVal(ctx context.Context, key string) (string, error)
	LRange(ctx context.Context, key string, from int, to int) ([]string, error)
	ZAdd(ctx context.Context, key string, score float64, member string) error
	ZRange(ctx context.Context, key string, start int64, stop int64) ([]string, error)
	ZRemRangeByRank(ctx context.Context, key string, start int64, stop int64) error
	ZIncrBy(ctx context.Context, key string, increment float64, member string) error
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error)
	ZRevRank(ctx context.Context, key string, member string) (int64, error)
	ZScore(ctx context.Context, key string, member string) (float64, error)

	GetSMembers(ctx context.Context, key string) ([]string, error)
	SetSAdd(ctx context.Context, key string, members ...interface{}) error
	SetNX(ctx context.Context, key string, seconds int64, data interface{}) (bool, error)
}
