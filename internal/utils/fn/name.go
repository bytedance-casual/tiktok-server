package fn

import (
	"reflect"
	"runtime"
	"strings"
)

func StructFuncName(i any) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	stringArr := strings.Split(fn, ".")
	fullName := stringArr[len(stringArr)-1]
	return strings.TrimSuffix(fullName, "-fm")
}

func FuncName(i any) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	stringArr := strings.Split(fn, ".")
	fullName := stringArr[len(stringArr)-1]
	return fullName
}
