package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ldtrieu/cerberus/package/errors"
)

// Cache structure
type Cache struct {
	rc        *redis.Client
	cacheTime time.Duration
}

// NewCache initializes Cache
func NewCache(rc *redis.Client, cacheTime time.Duration) ICache {
	return &Cache{
		rc:        rc,
		cacheTime: cacheTime,
	}
}

// Get reads value by key
func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := c.rc.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "get cache key=%s from redis failed", key)
	}
	return value, nil
}

// Set sets value by key
func (c *Cache) Set(ctx context.Context, key string, value []byte) error {
	err := c.rc.Set(ctx, key, value, c.cacheTime).Err()
	if err != nil {
		return errors.Wrap(err, "set value=%v with key=%s to redis failed", value, key)
	}
	return nil
}

func (c *Cache) SetObject(ctx context.Context, key string, val interface{}, duration time.Duration) error {
	dataBytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return c.rc.Set(ctx, key, dataBytes, duration).Err()
}

// Delete ...
func (c *Cache) Delete(ctx context.Context, key string) error {
	err := c.rc.Del(ctx, key).Err()
	if err != nil {
		return errors.Wrap(err, "delete with key=%s to redis failed", key)
	}
	return nil
}

// SetWithDuration ...
func (c *Cache) SetWithDuration(ctx context.Context, key string, value []byte, duration time.Duration) error {
	err := c.rc.Set(ctx, key, value, duration).Err()
	if err != nil {
		return errors.Wrap(err, "set value=%v with key=%s and duration=%d to redis failed", value, key, duration)
	}
	return nil
}

// LGet ...
func (c *Cache) LGet(ctx context.Context, key string) ([]byte, error) {
	value, err := c.rc.LPop(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "get cache key=%s from redis failed", key)
	}
	return value, nil
}

// LSet ...
func (c *Cache) LSet(ctx context.Context, key string, val []byte) error {
	return c.rc.LPush(ctx, key, val).Err()
}

// RSet ...
func (c *Cache) RSet(ctx context.Context, key string, val []byte) error {
	return c.rc.RPush(ctx, key, val).Err()
}

// LLen ...
func (c *Cache) LLen(ctx context.Context, key string) (int64, error) {
	val, err := c.rc.LLen(ctx, key).Result()
	if err != nil {
		return -1, errors.Wrap(err, "get len of key=%s from redis failed", key)
	}
	return val, nil
}

// LList ...
func (c *Cache) LList(ctx context.Context, key string) ([]string, error) {
	vals, err := c.rc.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, errors.Wrap(err, "get datas of key=%s from redis failed", key)
	}
	return vals, nil
}

// Decr ...
func (c *Cache) Decr(ctx context.Context, key string) (int64, error) {
	val, err := c.rc.Decr(ctx, key).Result()
	if err != nil {
		return -1, errors.Wrap(err, "Decrby key=%s from redis failed", key)
	}
	return val, err
}

// Incr ...
func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	val, err := c.rc.Incr(ctx, key).Result()
	if err != nil {
		return -1, errors.Wrap(err, "IncrBy key=%s from redis failed", key)
	}
	return val, err
}

// DecrBy ...
func (c *Cache) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	val, err := c.rc.DecrBy(ctx, key, value).Result()
	if err != nil {
		return -1, errors.Wrap(err, "Decrby key=%s from redis failed", key)
	}
	return val, err
}

// IncrBy ...
func (c *Cache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	val, err := c.rc.IncrBy(ctx, key, value).Result()
	if err != nil {
		return -1, errors.Wrap(err, "IncrBy key=%s from redis failed", key)
	}
	return val, err
}

// LRange ...
func (c *Cache) LRange(ctx context.Context, key string, from int, to int) ([]string, error) {
	result, err := c.rc.LRange(ctx, key, int64(from), int64(to)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "lrange cache key=%s from redis failed", key)
	}
	return result, nil
}

// ZIncrBy ...
func (c *Cache) ZIncrBy(ctx context.Context, key string, increment float64, member string) error {
	err := c.rc.
		ZIncrBy(ctx, key, increment, member).
		Err()
	if err != nil {
		return errors.Wrap(err, "zincrby member=%v with key=%s and increment=%v to redis failed. Error: %v", member, key, increment, err)
	}

	return nil
}

// ZAdd ...
func (c *Cache) ZAdd(ctx context.Context, key string, score float64, member string) error {
	err := c.rc.ZAdd(ctx, key, &redis.Z{
		Score:  score,
		Member: member,
	}).Err()

	if err != nil {
		return errors.Wrap(err, "zadd member=%v with key=%s and score=%v to redis failed", member, key, score)
	}

	return nil
}

// ZRange ...
func (c *Cache) ZRange(ctx context.Context, key string, start int64, stop int64) ([]string, error) {
	result, err := c.rc.ZRange(ctx, key, start, stop).Result()

	if err != nil {
		return nil, errors.Wrap(err, "zrange with key=%s, start=%v and stop=%v to redis failed", key, start, stop)
	}

	return result, nil
}

// ZRevRangeWithScores ...
func (c *Cache) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	result, err := c.rc.ZRevRangeWithScores(ctx, key, start, stop).Result()

	if err != nil {
		return nil, errors.Wrap(err, "ZRevRangeWithScores with key=%s, start=%v and stop=%v to redis failed. Error: %v", key, start, stop, err)
	}

	return result, nil
}

// ZRevRank ...
func (c *Cache) ZRevRank(ctx context.Context, key string, member string) (int64, error) {
	result, err := c.rc.ZRevRank(ctx, key, member).Result()
	if err != nil {
		return -1, errors.Wrap(err, "ZRevRank with key=%s, member=%v. Error: %v", key, member, err)
	}

	return result, nil
}

// ZScore ...
func (c *Cache) ZScore(ctx context.Context, key string, member string) (float64, error) {
	result, err := c.rc.ZScore(ctx, key, member).Result()
	if err == redis.Nil {
		return -1, err
	}

	if err != nil {
		return -1, errors.Wrap(err, "ZRevRank with key=%s, member=%v. Error: %v", key, member, err)
	}

	return result, nil
}

// ZRemRangeByRank ...
func (c *Cache) ZRemRangeByRank(ctx context.Context, key string, start int64, stop int64) error {
	err := c.rc.ZRemRangeByRank(ctx, key, start, stop).Err()

	if err != nil {
		return errors.Wrap(err, "zremrangebyrank with key=%s, start=%v and stop=%v to redis failed", key, start, stop)
	}

	return nil
}

// SetVal set value by key
func (c *Cache) SetVal(ctx context.Context, key string, value string) error {
	err := c.rc.Set(ctx, key, value, c.cacheTime).Err()
	if err != nil {
		return errors.Wrap(err, "set value=%v with key=%s to redis failed", value, key)
	}
	return nil
}

// GetVal reads value by key
func (c *Cache) GetVal(ctx context.Context, key string) (string, error) {
	return c.rc.Get(ctx, key).Result()
}

func (c *Cache) SetExpireTime(ctx context.Context, key string, seconds int64) error {
	return c.rc.Expire(ctx, key, time.Second*time.Duration(seconds)).Err()
}

func (c *Cache) Exists(ctx context.Context, keys ...string) int64 {
	return c.rc.Exists(ctx, keys...).Val()
}

func (c *Cache) GetSMembers(ctx context.Context, key string) ([]string, error) {
	data, err := c.rc.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Cache) SetSAdd(ctx context.Context, key string, members ...interface{}) error {
	return c.rc.SAdd(ctx, key, members).Err()
}

func (c *Cache) SetNX(ctx context.Context, key string, seconds int64, data interface{}) (bool, error) {
	return c.rc.SetNX(ctx, key, data, time.Second*time.Duration(seconds)).Result()
}
