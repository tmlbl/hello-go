// intermeths
package main

import (
	"fmt"
)

func main() {
	s := "Hello Goodbye"
	var ptr *string = &s
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Address: %p\n", ptr)
	fmt.Printf("Flat Pointer: %v\n", *ptr)
}
