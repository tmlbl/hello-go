// Pointers
// No pointer arithmetic is allowed in Go
// However, pointers are a cheap way to pass
// a variable around by reference. A pointer
// only takes up 4 bytes on a 32bit machine and
// 8 on a 64bit machine. This is a memory-cheap
// way to pass a large variable to a function
package main

import (
	"fmt"
)

func main() {
	s := "Hello Goodbye"
	// The & symbol before a var evaluates to its
	// pointer's address
	var ptr *string = &s
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Address: %p\n", ptr)
	// Putting the * before the pointer value
	// "flattens" it into the value it points to
	fmt.Printf("Flat Pointer: %v\n", *ptr)
}

/* OUTPUT:

String: Hello Goodbye
Address: 0xc21000a150
Flat Pointer: Hello Goodbye

*/
