package main

import "fmt"

const iter = 10

func Sqrt(x float64) (z float64) {
	z = 1.
	for i := 0; i < iter; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
}
