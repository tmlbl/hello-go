// types
package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

const (
	pi   = 3.14
	lang = "Go"
)

func main() {
	var a, b, c rune = 34, 56, 91
	fmt.Println(a, b, c)
	var sal = Salutation{"Joe", "Sup"}
	fmt.Println(sal.greeting, ", ", sal.name, "!")
}
