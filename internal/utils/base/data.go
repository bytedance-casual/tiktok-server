package base

func DefaultOrValue[T int | int64 | string](data T, defaultValue T) T {
	var d T
	if data == d {
		return defaultValue
	}
	return data
}

func IsDefaultValue[T int | int64 | string](data T) bool {
	var d T
	if data == d {
		return true
	}
	return false
}
