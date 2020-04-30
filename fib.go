package main

import "fmt"

// fibonacci ...
func fibonacci() func() int {
	prev := 0
	curr := 0
	return func() int {
		res := curr
		if curr == 0 {
			curr = 1
		} else {
			curr = prev + curr
			prev = res
		}
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
