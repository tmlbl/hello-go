// Funcs Reloaded
// Functions can return other functions
// Endless functional programming possibilities
package main

import (
	"fmt"
)

func main() {
	square := powerOf(2)
	cube := powerOf(3)
	n := 3
	i := 4
	fmt.Printf("%d squared is %d\n", n, square(n))
	fmt.Printf("%d squared is %d\n", i, square(i))
	fmt.Printf("%d cubed is %d\n", n, cube(n))
	fmt.Printf("%d cubed is %d\n", i, cube(i))
}

// A function factory that takes a power value
// and returns a function that will raise an int
// to that power
func powerOf(pow int) func(int) int {
	return func(num int) int {
		res := 1
		for i := 0; i < pow; i++ {
			res *= num
		}
		return res
	}
}

/* OUTPUT:

3 squared is 9
4 squared is 16
3 cubed is 81
4 cubed is 256

*/
