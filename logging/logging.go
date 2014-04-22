// Logging
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("%s\n", thing())
	// The shortfile flag will print the
	// filename and line number with every
	// log message
	log.SetFlags(log.Lshortfile)
	// A function can be timed by simply
	// setting start and end timestamps
	// then printing the difference
	start := time.Now()
	thing()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Function took: %s\n", delta)
	// Log has the same print functions as fmt,
	// and they work in mostly the same way
	// This empty print will only print the
	// shortfile flag
	log.Print("\n")
}

func thing() string {
	time.Sleep(10000)
	return "hello woooorld"
}
