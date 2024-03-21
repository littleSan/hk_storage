package redisCli

import (
	"github.com/go-redis/redis"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
	"strings"
	"time"
)

var _ Repo = (*cacheRepo)(nil)

type Repo interface {
	i()
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
	TTL(key string) (time.Duration, error)
	Expire(key string, ttl time.Duration) bool
	ExpireAt(key string, ttl time.Time) bool
	Del(key string) bool
	Exists(keys ...string) bool
	Incr(key string) int64
	Close() error
	Version() string
	LimitCount(keys string, maxTodayCount int) (bool, error)
	LimitAmount(keys string) (int, error)
}

type cacheRepo struct {
	client *redis.Client
}

func New() Repo {
	//cc, _ = redisConnect()
	return &cacheRepo{
		client: cc,
	}
}

func (i *cacheRepo) i() {
}

var cc *redis.Client

func RegisterRedis() (*redis.Client, error) {
start:
	if cc != nil {
		cc.Close()
		cc = nil
	}

	cc = redis.NewClient(&redis.Options{
		Addr:     configs.TomlConfig.Redis.Addr,
		Password: configs.TomlConfig.Redis.Password,
		DB:       0,
	})

	if err := cc.Ping().Err(); err != nil {
		logger.Info("ping err :", err)
		time.Sleep(5 * time.Second)
		goto start
	}

	return cc, nil
}

// Set set some <key,value> into redis
func (c *cacheRepo) Set(key, value string, ttl time.Duration) error {

	if err := c.client.Set(key, value, ttl).Err(); err != nil {
		logger.Trace("redis set key: %s err", key)
		return err
	}

	return nil
}

// Get get some key from redis
func (c *cacheRepo) Get(key string) (string, error) {

	value, err := c.client.Get(key).Result()
	if err != nil {
		logger.Trace("redis get key: %s err", key)
		return "", err
	}

	return value, nil
}

// TTL get some key from redis
func (c *cacheRepo) TTL(key string) (time.Duration, error) {
	ttl, err := c.client.TTL(key).Result()
	if err != nil {
		return -1, err
	}

	return ttl, nil
}

// Expire expire some key
func (c *cacheRepo) Expire(key string, ttl time.Duration) bool {
	ok, _ := c.client.Expire(key, ttl).Result()
	return ok
}

// ExpireAt expire some key at some time
func (c *cacheRepo) ExpireAt(key string, ttl time.Time) bool {
	ok, _ := c.client.ExpireAt(key, ttl).Result()
	return ok
}

func (c *cacheRepo) Exists(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := c.client.Exists(keys...).Result()
	return value > 0
}

func (c *cacheRepo) Del(key string) bool {

	if key == "" {
		return true
	}

	value, _ := c.client.Del(key).Result()
	return value > 0
}

func (c *cacheRepo) Incr(key string) int64 {

	value, _ := c.client.Incr(key).Result()
	return value
}

// Close close redis client
func (c *cacheRepo) Close() error {
	return c.client.Close()
}

// Version redis server version
func (c *cacheRepo) Version() string {
	server := c.client.Info("server").Val()
	spl1 := strings.Split(server, "# Server")
	spl2 := strings.Split(spl1[1], "redis_version:")
	spl3 := strings.Split(spl2[1], "redis_git_sha1:")
	return spl3[0]
}

// 限制单位时间内的次数
func (c *cacheRepo) LimitCount(keys string, maxTodayCount int) (bool, error) {

	key := "hk_storage:" + keys + ":" + time.Now().Format("2006-01-02")
	val, err := c.client.Get(key).Int()
	if err != nil && err != redis.Nil {
		return false, err
	}

	if err == redis.Nil {
		_, err := c.client.Set(key, 1, 24*time.Hour).Result()
		if err != nil {
			return false, err
		}
	} else {
		if val >= maxTodayCount {
			return false, nil // 今天已达到最大登录次数
		}
		_, err := c.client.Incr(key).Result()
		if err != nil {
			return false, err
		}
	}

	return true, nil // 未达到限制次数
}

// 限制单位时间内的次数
func (c *cacheRepo) LimitAmount(keys string) (int, error) {

	key := "hk_storage:" + keys + ":" + time.Now().Format("2006-01-02")
	val, err := c.client.Get(key).Int()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	return val, nil // 未达到限制次数
}
