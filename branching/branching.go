// Branching
// Go offers three kinds of control structures:
// * if/else statements
// * switch/case statements
// * select construct
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand must be seeded with a non-fixed val
	// UnixNano() converts the unix timestamp to
	// an int64
	rand.Seed(time.Now().UnixNano())
	// random integer between 1 and 10
	val := rand.Intn(10)
	if val < 5 { // Note the lack of parens
		fmt.Printf("%d is less than 5\n", val)
	} else if val > 5 {
		fmt.Printf("%d is greater than 5\n", val)
	} else {
		fmt.Printf("%d is equal to 5\n", val)
	}
}
