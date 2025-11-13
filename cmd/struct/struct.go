package main

import "cmp"

type CustomData interface {
	cmp.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}
