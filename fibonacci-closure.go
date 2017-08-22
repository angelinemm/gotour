package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    current := 0 
	next := 1
	return func() int {
	    to_return := current
		current = next
		next = next + to_return
		return to_return 
	}	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
