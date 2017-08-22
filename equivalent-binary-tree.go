package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// RecursWalk the tree t sending all values
// from the tree to the channel ch.
func RecursWalk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
	    RecursWalk(t.Left, ch)
	}
	ch <- t.Value
    if t.Right != nil {
	    RecursWalk(t.Right, ch)
	}
}

// Walk the tree from the root sending all values
// from the tree to the channel
// and then closing the channel
func Walk(t *tree.Tree, ch chan int) {
    RecursWalk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1 := make(chan int, 10)
	go Walk(t1, c1)
	
	c2 := make(chan int, 10)
	go Walk(t2, c2)
	
	for i1 := range(c1) {
	    select {
		    case i2 := <- c2:
			    if i1 != i2 {
				    return false
				}
			default:
			    break
		}
	}
	return true
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
