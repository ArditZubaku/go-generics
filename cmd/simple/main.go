package main

import "cmp"

func Add1[T int | float64](a T, b T) T {
	return a + b
}

type Num interface {
	int | int8 | int16 | int64 | float32 | float64
}

func Add2[T Num](a T, b T) T {
	return a + b
}

func Add3[T cmp.Ordered](a T, b T) T {
	return a + b
}

type UserID int

// Add4 - tilda allows to be used any type that is an alias of the underlying type
func Add4[T ~int | float64](a T, b T) T {
	return a + b
}

func main() {
	result := Add1(1, 2)
	println(result)

	a := UserID(1)
	b := UserID(2)
	// println(Add1(a, b))
	println(Add4(a, b))
}
