// Funcs Reloaded
// Functions can return other functions
// Endless functional programming possibilities
package main

import (
	"fmt"
)

func main() {
	square := powerOf(2)
	n := 3
	fmt.Printf("%d squared is %d\n", n, square(n))
	cube := powerOf(3)
	fmt.Printf("%d cubed is %d\n", n, cube(n))
}

// A function factory that takes a power value
// and returns a function that will raise an int
// to that power
func powerOf(pow int) (res func(int) int) {
	return func(num int) int {
		for i := 1; i < pow; i++ {
			num *= num
		}
		return num
	}
}
