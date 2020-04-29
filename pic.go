package main

import "golang.org/x/tour/pic"

// Produces a slice of length dy where each element is a
// slice of length dx each integer represents a pixel in
// the resulting image.
func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for i := range res {
		res[i] = make([]uint8, dx)
		for j := range res[i] {
			res[i][j] = uint8((i + j) / 2)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
