// FizzBuzz
// It's about that time, y'all
package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		f := ""
		if i%3 == 0 {
			f += "Fizz"
		}
		if i%5 == 0 {
			f += "Buzz"
		}
		if i == 0 {
			f = ""
		}
		fmt.Printf("%d %s\n", i, f)
	}
}
