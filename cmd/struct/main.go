package main

func main() {
	u := User[int]{
		ID:   0,
		Name: "",
		Data: 3,
	}

	println("User: ", u)

	u2 := User[string]{
		ID:   0,
		Name: "",
		Data: "3",
	}

	println("User: ", u2)
}
