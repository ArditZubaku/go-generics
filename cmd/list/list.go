package main

// func Has(list []string, value string) bool {
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
//
// 	return false
// }

// func Has[T comparable](list []T, value T) bool {
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
//
// 	return false
// }

import "slices"

func Has[T comparable](list []T, value T) bool {
	return slices.Contains(list, value)
}

func NewEmptyList[T any]() []T {
	return make([]T, 0)
}
