package main

func main() {
	println("Has a: ", Has([]string{"a", "b"}, "a"))
	println("Has c: ", Has([]string{"a", "b"}, "c"))
	println("Has 2: ", Has([]int{1, 2, 3}, 2))

	println("MyList: ", NewEmptyList[int]())
}
