package main

func main() {
	println("Has 2: ", HasAny([]int{1, 2, 3}, 2, func(a, b int) bool {
		return a == b
	}))
	println("Has 2: ", HasComparable([]int{1, 2, 3}, 2))
	println("Has 2: ", HasMyInterface([]ID{1, 2, 3}, 2))
	println("Has 2: ", HasInt([]int{1, 2, 3}, 2))
	println("Has 2: ", HasInt([]ID{1, 2, 3}, 2))
	PrintBalance(Euro(250))
}
