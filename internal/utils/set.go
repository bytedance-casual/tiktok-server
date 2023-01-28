package utils

import (
	"sync"
)

type void struct{} //空结构不占用任何内存

// Set 基于sync.Map实现的Set
type Set[T string] struct {
	members *sync.Map
}

func NewSet[T string]() *Set[T] {
	return &Set[T]{
		members: new(sync.Map),
	}
}

func (set *Set[T]) Contains(element T) bool {
	_, ok := set.members.Load(element)
	if ok {
		return true
	}
	return false
}

func (set *Set[T]) Poll() (T, bool) {
	if set.Empty() {
		return *new(T), false
	}
	var result T
	set.members.Range(func(key, value interface{}) bool {
		element := key.(T)
		if set.Remove(element) {
			result = element
			return false
		}
		return true
	})
	return result, true
}

func (set *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		set.members.Store(element, struct{}{})
	}
}

func (set *Set[T]) Remove(element T) bool {
	if !set.Contains(element) {
		return false //不包含，删除不成功
	}
	set.members.Delete(element)
	return true
}

func (set *Set[T]) All() []*T {
	result := make([]*T, 0)
	set.members.Range(func(key, value interface{}) bool {
		element := key.(T)
		result = append(result, &element)
		return true
	})
	return result
}

func (set *Set[T]) Empty() bool {
	empty := true
	set.members.Range(func(key, value interface{}) bool {
		empty = false
		return false
	})
	return empty
}
