// Loops
package main

import (
	"fmt"
)

func main() {
	message := "Grappa Dappa Doo"
	// For: range vars are key(or index) and value
	// If the index is not needed, it must be replaced
	// with the _ (discard) character, or it will throw
	// an unused variable error
	for _, s := range message {
		fmt.Println("Rune code ", s)
		fmt.Printf("Letter: %c\n", s)
	}
}
