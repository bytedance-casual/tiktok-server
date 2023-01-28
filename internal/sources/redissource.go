package sources

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"tiktok-server/internal/model"
	"time"
)

var RedisSource = &redisSource{}

func init() {
	Context.RegisterSource(RedisSource)
}

type redisSource struct {
	Client *redis.Client
}

func (source *redisSource) ConfigAvailable(config *model.Config) bool {
	if config.Redis != nil {
		return true
	}
	return false
}

func (source *redisSource) Setup(config *model.Config) {
	source.Connect(config)
}

func (source *redisSource) Connect(config *model.Config) {
	conf := config.Redis
	if conf == nil {
		return
	}
	options := &redis.Options{}
	err := copier.Copy(options, conf)
	if err != nil {
		log.Panicf("[redisdata.Setup]:%s", err)
	}
	source.Client = redis.NewClient(options)
	//测试连接
	_, err = source.Client.Ping().Result()
	if err != nil {
		log.Panicf("[redisdata.Conect]:%s", err)
	}
	log.Printf("[redisdata.Setup]:Setup successfully:%v", conf)
}

// HSet 向哈希表中插入字符串
func (source *redisSource) HSet(key string, field string, value any) error {
	if err := source.Client.HSet(key, field, value).Err(); err != nil {
		return errors.Wrap(err, "HSet cache key ==========> "+key)
	}
	return nil
}

// HGet 仅获取哈希表中值的字符串
func (source *redisSource) HGet(key string, field string) (string, error) {
	str, err := source.Client.HGet(key, field).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", errors.Wrap(err, "HGet cache key ==========> "+key)
	}
	return str, nil
}

func (source *redisSource) HDel(key string, field string) error {
	err := source.Client.HDel(key, field).Err()
	if err != nil {
		return errors.Wrap(err, "HDel cache key ==========> "+key)
	}
	return nil
}

func (source *redisSource) Get(key string) (string, error) {
	str, err := source.Client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", errors.Wrap(err, "Get cache key ==========> "+key)
	}
	return str, nil
}

func (source *redisSource) GetDel(key string) (string, error) {
	str, err := source.Client.Do("getdel", key).String()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", errors.Wrap(err, "Get cache key ==========> "+key)
	}
	return str, nil
}

func (source *redisSource) SetOne(key string, value string, expSec time.Duration) error {
	if err := source.Client.Set(key, value, expSec*time.Second).Err(); err != nil {
		return errors.Wrap(err, "Set To Cache ==========> "+key)
	}
	return nil
}

func (source *redisSource) SetMany(values map[string]any, exp time.Duration) error {
	kvs := make([]any, len(values)*2)
	i := 0
	for k, v := range values {
		kvs[i], kvs[i+1] = k, v
		i += 2
	}
	if err := source.Client.MSet(kvs...).Err(); err != nil {
		return errors.Wrap(err, " Set Many To Cache ==========> "+strconv.Itoa(len(values)))
	}
	return nil
}

func (source *redisSource) Del(keys ...string) error {
	if err := source.Client.Del(keys...).Err(); err != nil {
		return errors.Wrap(err, "==========> Delete ")
	}
	return nil
}

func (source *redisSource) Close() {
	err := source.Client.Close()
	if err != nil {
		log.Printf("[redisSource.Close]:")
	}
}