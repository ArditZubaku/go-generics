package main

// HasInt - Must be of that type or underlying type
func HasInt[T ~int](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

// func HasInt[T int](list []T, value T) bool {
//
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
//
// 	return false
// }
