package utils

func Ternary[T string](condition bool, trueData T, falseData T) T {
	if condition {
		return trueData
	}
	return falseData
}
