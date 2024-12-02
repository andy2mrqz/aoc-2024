package utils

import "strconv"

// Map takes the slice s and applies the function fn to each element, returning
// a slice whose type corresponds to the return type of fn
func Map[T, V any](s []T, fn func(T) V) []V {
	out := make([]V, len(s))
	for i, item := range s {
		out[i] = fn(item)
	}
	return out
}

// StrToInt takes the string s and returns it as an int, panicking if invalid
func StrToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

// AbsInt takes an int i and returns its absolute value, along with its original sign (-1, 0, 1)
func AbsInt(i int) (int, int) {
	if i < 0 {
		return -i, -1
	} else if i == 0 {
		return i, 0
	}
	return i, 1
}
