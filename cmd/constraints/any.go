package main

// func Has[T any](list []T, value T) bool {
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
//
// 	return false
// }

func HasAny[T any](list []T, value T, equalFunc func(a, b T) bool) bool {
	for _, v := range list {
		if equalFunc(v, value) {
			return true
		}
	}

	return false
}
