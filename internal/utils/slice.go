package utils

import (
	"encoding/json"
	"errors"
	"sort"
)

// DoubleSlice 将数组的所有元素复制一份并加在源数组后并返回
// 如果是指针类型复制的是指针
func DoubleSlice[T any](origin []T) []T {
	newSlice := make([]T, len(origin))
	copy(newSlice, origin)
	origin = append(origin, newSlice...)
	return origin
}

func Find[T Comparator](slice []T, val T) (int, bool) {
	for i, item := range slice {
		if item.Equals(val) {
			return i, true
		}
	}
	return -1, false
}

func FindBasic[T int | string](slice []T, val T) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Remove[T Comparator](slice []T, val T) ([]T, error) {
	flag := false
	for i, item := range slice {
		if item.Equals(val) {
			flag = true
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	if flag != true {
		return nil, errors.New("cannot find element")
	}
	return slice, nil
}

type Comparator interface {
	Equals(o2 interface{}) bool
}

func ToString[T interface{}](slice []T) (string, error) {
	marshal, err := json.Marshal(slice)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

func Rebuild[T string | int | int64](sliceStr string) ([]T, error) {
	var result []T
	err := json.Unmarshal([]byte(sliceStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ConditionFind 根据给定条件查找
// 条件使用一个返回bool值的函数给出
// 如果没有找到，直接返回nil
func ConditionFind[T interface{}](slice []*T, condition func(index int) bool) *T {
	if slice == nil || len(slice) == 0 {
		return nil
	}
	for i := 0; i < len(slice); i++ {
		if condition(i) {
			return slice[i]
		}
	}
	return nil
}

type sortSlice[T interface{}] struct {
	Slice []T
	Fun   func(i, j int, slice []T) bool
}

func newSortSlice[T interface{}](slice []T, fun func(i, j int, slice []T) bool) *sortSlice[T] {
	return &sortSlice[T]{Slice: slice, Fun: fun}
}

func SortAscending[T interface{}](slice []T, lessFunc func(i, j int, slice []T) bool) { //升序排列
	s := newSortSlice[T](slice, lessFunc)
	sort.Sort(s)
}

func SortDescending[T interface{}](slice []T, lessFunc func(i, j int, slice []T) bool) { //降序排列
	s := newSortSlice[T](slice, lessFunc)
	sort.Sort(sort.Reverse(s))
}

func (t sortSlice[T]) Len() int {
	return len(t.Slice)
}

func (t sortSlice[T]) Less(i, j int) bool {
	return t.Fun(i, j, t.Slice)
}

func (t sortSlice[T]) Swap(i, j int) {
	t.Slice[i], t.Slice[j] = t.Slice[j], t.Slice[i]
}
