// Funcs
// Functions in Go are defined using the func keyword.
// Like JS, they can be passed around as parameters to
// other functions. They can also return multiple typed
// values. This creates the opportunity to return error
// codes without overriding the function's normal return
package main

import (
	"fmt"
)

// instead of classes, define custom types
type Salutation struct {
	name     string
	greeting string
}

// takes print func as a param
func Greet(sal Salutation, p func(string)) {
	var greeter string = sal.greeting + ", " + sal.name + "!"
	p(greeter)
}

func PrintString(s string) {
	fmt.Println(s)
}

// can return multiple values from a func
// can name return types and return by assignment
func CountLetters(sal Salutation) (nameLen int, greetLen int) {
	// len is a standalone function, not a method
	nameLen = len(sal.name)
	greetLen = len(sal.greeting)
	return
}

func main() {
	var sal = Salutation{}
	sal.name = "Billy Bob"
	sal.greeting = "Good Evening"
	// pass func PrintString as a param
	Greet(sal, PrintString)
	var nameLen, greetLen = CountLetters(sal)
	fmt.Println(nameLen, greetLen)
}
