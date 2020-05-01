package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walkHelper func(t *tree.Tree)
	walkHelper = func(t *tree.Tree) {
		if t != nil {
			walkHelper(t.Left)
			ch <- t.Value
			walkHelper(t.Right)
		}
	}
	walkHelper(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var same bool = true
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			same = same && (v1 == v2)
		} else if ok1 || ok2 {
			return false
		} else {
			break
		}
	}
	return same
}

func main() {
	s := 20000
	fmt.Println(Same(tree.New(2), tree.New(3)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(200), tree.New(200)))
	fmt.Println(Same(tree.New(s), tree.New(s)))
}
