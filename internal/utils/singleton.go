package utils

import "sync"

type singleton interface{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
