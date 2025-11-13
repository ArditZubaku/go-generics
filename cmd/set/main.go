package main

func main() {
	intSet := NewSet(1, 2, 3, 1, 4)
	println("Has 2: ", intSet.Has(2))
	println("Has 5: ", intSet.Has(5))

	strSet := NewSet("a", "b")
	println("Has a: ", strSet.Has("a"))
	println("Has c: ", strSet.Has("c"))
}
