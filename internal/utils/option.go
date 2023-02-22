package utils

func Ternary[T string | bool | int64](condition bool, trueData T, falseData T) T {
	if condition {
		return trueData
	}
	return falseData
}
