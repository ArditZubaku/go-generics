package main

import "cmp"

// input: [1,2,3], (n) => n * 2
// output: [2,3,6]

func MapValues(values []int, mapFunc func(int, int) int) []int {
	var newValues []int
	for i, v := range values {
		newValues = append(newValues, mapFunc(i, v))
	}

	return newValues
}

func MapValuesGeneric[T cmp.Ordered](values []T, mapFunc func(int, T) T) []T {
	var newValues []T
	for i, v := range values {
		newValues = append(newValues, mapFunc(i, v))
	}

	return newValues
}

type CustomMap[T comparable, V int | string] map[T]V
