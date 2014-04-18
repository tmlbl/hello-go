/**
 * My awesome library
 * behold
 */

package main

import (
	"fmt"
)

func main() {
	truthy := true
	message := "Grappa Dappa Doo"
	if truthy {
		fmt.Println(message)
	}
	for _, s := range message {
		fmt.Println("Rune code ", s)
	}
}
