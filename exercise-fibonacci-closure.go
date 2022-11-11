package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n0, n1 := 0, 1
	return func() int {
		ret := n0
		n0, n1 = n1, n0+n1
		return ret
	}
}


func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

