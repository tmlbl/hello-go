// Types
// Go does not have OO-style objects, but they
// can be approximated using structs and custom
// types. A new type can be created by the programmer
// and then used in the same way as a built-in type
package main

import (
	"fmt"
)

// Defining a custom type
type Salutation struct {
	name     string
	greeting string
}

// Method for declaring immutable constants. Note that
// whereas a declared but unused variable will throw an
// error at compile time, constants have no such
// restrictions
const (
	pi   = 3.14
	lang = "Go"
)

func main() {
	// Initialization of a custom type, similar to invoking
	// an object constructor. Note the curly braces instead
	// of parenthesis
	var sal = Salutation{"Joe", "Sup"}
	fmt.Println(sal.greeting, ", ", sal.name, "!")
}
