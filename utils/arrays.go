package utils

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

type Reducer[T, V any] func(accum T, value V) T

func Reduce[T, V any](xs []V, f Reducer[T, V]) T {
	var accum T
	for _, x := range xs {
		accum = f(accum, x)
	}
	return accum
}

func Matrix[T any](rows int, cols int) [][]T {
	m := make([][]T, rows)
	for i := range m {
		m[i] = make([]T, cols)
	}
	return m
}

func SliceIndex[T comparable](ts []T, elem T) int {
	for i := 0; i < len(ts); i++ {
		if ts[i] == elem {
			return i
		}
	}
	return -1
}
