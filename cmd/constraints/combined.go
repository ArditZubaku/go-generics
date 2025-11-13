package main

import (
	"fmt"
	"math"
)

type Currency interface {
	~int | ~int64
	ISO4127Code() string
	Decimal() int
}

func PrintBalance[T Currency](b T) {
	balance := float64(b) / math.Pow10(b.Decimal())
	fmt.Printf("%.*f %s\n", b.Decimal(), balance, b.ISO4127Code())
}

type Euro int64

func (e Euro) ISO4127Code() string {
	return "EUR"
}

func (e Euro) Decimal() int {
	return 2
}
