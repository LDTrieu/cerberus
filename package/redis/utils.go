package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func Get(redisClient *redis.Client, ctx context.Context, key string) (string, error) {
	dataBytes, err := redisClient.Get(ctx, key).Bytes()
	if err != nil && err != redis.Nil {
		return "", err
	}

	return string(dataBytes), nil
}

func Set(redisClient *redis.Client, ctx context.Context, key string, seconds int64, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, dataBytes, time.Second*time.Duration(seconds)).Err()
}

func Delete(redisClient *redis.Client, ctx context.Context, keys ...string) error {
	return redisClient.Del(ctx, keys...).Err()
}

func GetHash(redisClient *redis.Client, ctx context.Context, key string) (interface{}, error) {
	var data interface{}
	err := redisClient.HGetAll(ctx, key).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SetHash(redisClient *redis.Client, ctx context.Context, key string, seconds int64, data interface{}) error {
	return redisClient.HSet(ctx, key, data, time.Second*time.Duration(seconds)).Err()
}

func GetSMembers(redisClient *redis.Client, ctx context.Context, key string) ([]string, error) {
	data, err := redisClient.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}
func IsSetMember(redisClient *redis.Client, ctx context.Context, key string, member string) (bool, error) {
	is, err := redisClient.SIsMember(ctx, key, member).Result()
	if err != nil {
		return false, err
	}
	return is, nil
}
func IsSetMembers(redisClient *redis.Client, ctx context.Context, key string, members ...interface{}) (bool, error) {
	is, err := redisClient.SMIsMember(ctx, key, members...).Result()
	if err != nil {
		return false, err
	}
	for _, v := range is {
		if v {
			return true, nil
		}
	}
	return false, nil
}
func SetSAdd(redisClient *redis.Client, ctx context.Context, key string, members ...interface{}) error {
	return redisClient.SAdd(ctx, key, members).Err()
}
func Exists(redisClient *redis.Client, ctx context.Context, keys ...string) int64 {
	count := redisClient.Exists(ctx, keys...).Val()
	return count
}
