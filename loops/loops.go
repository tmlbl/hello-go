// Loops
// There is only one looping keyword in Go: for
// However, it can execute any common loop
package main

import (
	"fmt"
)

func main() {
	// You can write a traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration #%d\n", i+1)
	}
	// Example of a for loop iterating a string
	message := "Grappa"
	// For: range vars are key(or index) and value
	// If the index is not needed, it must be replaced
	// with the _ (discard) character, or it will throw
	// an unused variable error
	for _, s := range message {
		fmt.Println("Rune code ", s)
		fmt.Printf("Letter: %c\n", s)
	}
}

/* OUTPUT:

Iteration #1
Iteration #2
Iteration #3
Iteration #4
Iteration #5
Rune code  71
Letter: G
Rune code  114
Letter: r
Rune code  97
Letter: a
Rune code  112
Letter: p
Rune code  112
Letter: p
Rune code  97
Letter: a

*/
