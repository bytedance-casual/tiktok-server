package repo

import (
	"encoding/json"
	"github.com/pkg/errors"
	"tiktok-server/internal/sources"
	"time"
)

type CacheOnlyRepo[D interface{}] struct{}

func (rep *CacheOnlyRepo[D]) SetupRepo() {}

// FindOne 从redis中获取model:key的数据并反序列化
func (rep *CacheOnlyRepo[D]) FindOne(key string) (*D, error) {
	str, err := sources.RedisSource.Get(key)
	if str == "" {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	m := new(D)
	err = json.Unmarshal([]byte(str), m)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal Data From Cache ==========> ")
	}
	return m, err
}

// SaveOne  序列化结构体并保存到redis
func (rep *CacheOnlyRepo[D]) SaveOne(key string, data *D, expSecond time.Duration) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "Unmarshal Data To Cache ==========> ")
	}
	return sources.RedisSource.SetOne(key, string(bytes), expSecond*time.Second)
}

func (rep *CacheOnlyRepo[D]) DeleteOne(key string) error {
	err := sources.RedisSource.Del(key)
	if err != nil {
		return errors.Wrap(err, "Del Key ==========> "+key)
	}
	return nil
}
