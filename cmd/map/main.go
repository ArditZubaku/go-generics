package main

func main() {
	mappedValues := MapValues(
		[]int{1, 2, 3},
		func(_, val int) int {
			return val * 2
		},
	)

	println("Mapped values: ", mappedValues)

	mappedValuesFloat64 := MapValuesGeneric(
		[]float64{1, 2, 3},
		func(_ int, val float64) float64 {
			return val * 2
		},
	)

	println("Mapped values: ", mappedValuesFloat64)

	mappedValuesFloat32 := MapValuesGeneric(
		[]float32{1, 2, 3},
		func(_ int, val float32) float32 {
			return val * 2
		},
	)

	println("Mapped values: ", mappedValuesFloat32)

	m := make(CustomMap[int, string])
	m[3] = "3"
	println("CustomMap: ", m)
}
