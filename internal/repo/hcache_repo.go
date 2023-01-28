package repo

import (
	"github.com/pkg/errors"
	"reflect"
	"tiktok-server/internal/sources"
	"tiktok-server/internal/utils"
)

type HCacheRepo[V interface{}] struct {
	MapName string //key为结构体字符串
}

func (rep *HCacheRepo[V]) SetupRepo() {
	var v V
	rep.MapName = reflect.ValueOf(v).Type().Name()
}

func (rep *HCacheRepo[V]) Set(field string, v V) error {
	kStr := field
	vStr := utils.Serialize[V](v)
	err := sources.RedisSource.HSet(rep.MapName, kStr, vStr)
	if err != nil {
		return err
	}
	return nil
}

func (rep *HCacheRepo[V]) Get(field string) (*V, error) {
	value, err := sources.RedisSource.HGet(rep.MapName, field)
	if err != nil {
		return nil, err
	}
	if value == "" {
		return nil, nil
	}
	return utils.Deserialize[V](value), nil
}

func (rep *HCacheRepo[V]) DeleteField(field string) error {
	if sources.RedisSource.Client.HExists(rep.MapName, field).Val() {
		err := sources.RedisSource.HDel(rep.MapName, field)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rep *HCacheRepo[V]) GetAll() (map[string]*V, error) {
	result, err := sources.RedisSource.Client.HGetAll(rep.MapName).Result()
	data := make(map[string]*V)
	if err != nil {
		if err.Error() == "redis: nil" {
			return data, nil
		}
		return nil, errors.Wrap(err, "Get All Data From RedisHash")
	}

	for field, value := range result {
		r := utils.Deserialize[V](value)
		data[field] = r
	}
	return data, nil
}
