// Maps - like a hash map
package main

import (
	"fmt"
)

func main() {
	myMap := map[string]string{
		"don":  "johnson",
		"doop": "moog",
	}
	myMap["foo"] = "bar"
	fmt.Println(myMap)
	for k, v := range myMap {
		fmt.Println(k, v)
		if k == "foo" {
			fmt.Println("Aw shit it's foo")
		}
	}
}
