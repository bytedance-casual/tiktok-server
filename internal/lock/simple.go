package lock

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"tiktok-server/internal/sources"
	"tiktok-server/internal/utils"
	"time"
)

const (
	DefaultAutoUnlockTime = 1 * time.Second
	Prefix                = "L:"
)

// region 基于redis的原始版本，可能会有误删问题

type SimpleLock struct {
	key       string
	duration  time.Duration
	expiredAt time.Time
}

func NewSimpleLock(key string, duration ...time.Duration) *SimpleLock {
	var expire time.Duration
	if duration == nil || len(duration) == 0 {
		expire = DefaultAutoUnlockTime
	} else {
		expire = duration[0]
	}
	return &SimpleLock{key: key, expiredAt: time.Now().Add(expire), duration: expire}
}

func (lock *SimpleLock) TryLock() (bool, error) {
	v := time.Now().String()
	result, err := sources.RedisSource.Client.SetNX(Prefix+lock.key, v, lock.duration).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (lock *SimpleLock) UnLock() {
	_, _ = sources.RedisSource.Client.Del(lock.key).Result()
}

// endregion

// region 解决误删问题的版本，但是高并发情况下，获取锁与释放锁的原子性得不到保证

func init() {
	rand.Seed(time.Now().UnixNano())
}

type WithOwnerJudgeLock struct {
	key       string
	duration  time.Duration
	expiredAt time.Time
	owner     string
}

func NewWithOwnerJudgeLock(key string, duration ...time.Duration) *WithOwnerJudgeLock {
	var expire time.Duration
	if duration == nil || len(duration) == 0 {
		expire = DefaultAutoUnlockTime
	} else {
		expire = duration[0]
	}
	return &WithOwnerJudgeLock{
		key:       key,
		expiredAt: time.Now().Add(expire),
		duration:  expire,
		owner:     utils.RandomString(10),
	}
}

func (lock *WithOwnerJudgeLock) TryLock() (bool, error) {
	v := lock.owner
	result, err := sources.RedisSource.Client.SetNX(Prefix+lock.key, v, lock.duration).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (lock *WithOwnerJudgeLock) UnLock() error {
	//判断是否为锁的持有者
	v, err := sources.RedisSource.Get(Prefix + lock.key)
	if err != nil {
		return errors.Wrap(err, "unlock: "+lock.key)
	}
	if v == lock.owner {
		//释放自己持有的锁
		_, _ = sources.RedisSource.Client.Del(lock.key).Result()
	} else {
		return errors.Wrap(err, fmt.Sprintf(" %s try unlock other's lock: %s", lock.owner, v))
	}
	return nil
}

// endregion
