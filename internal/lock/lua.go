package lock

import (
	"github.com/go-redis/redis"
	"github.com/smallnest/rpcx/log"
	"strconv"
	"tiktok-server/internal/sources"
	"tiktok-server/internal/utils"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//重复上锁的时候，将锁的过期时间续期
	lockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`
	unlockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`
	randomLen = 16
	// 默认超时时间，防止死锁
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
)

type RedisLock struct {
	key       string
	duration  time.Duration
	expiredAt time.Time
	owner     string
}

func NewRedisLock(key string, duration ...time.Duration) *RedisLock {
	var expire time.Duration
	if duration == nil || len(duration) == 0 {
		expire = DefaultAutoUnlockTime
	} else {
		expire = duration[0]
	}
	return &RedisLock{
		key:       key,
		expiredAt: time.Now().Add(expire),
		duration:  expire,
		owner:     utils.RandomString(10),
	}
}

func (lock *RedisLock) TryLock() (bool, error) {
	resp, err := sources.RedisSource.Client.Eval(lockCommand, []string{Prefix + lock.key}, []string{
		lock.owner, strconv.FormatInt(lock.duration.Milliseconds(), 10)}).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		log.Errorf("Error on acquiring lock for %s, %s", lock.key, err.Error())
		return false, err
	} else if resp == nil {
		return false, nil
	}
	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	}
	return false, nil
}

func (lock *RedisLock) UnLock() (bool, error) {
	resp, err := sources.RedisSource.Client.Eval(unlockCommand, []string{Prefix + lock.key}, []string{lock.owner}).Result()
	if err != nil {
		return false, err
	}
	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}
	return reply == 1, nil
}
