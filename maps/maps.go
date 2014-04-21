// Maps
// Allows for key-value pairs to be stored in an
// array-like construct. Maps can be extended after
// they are declared, but have strong types for keys
// and values
package main

import (
	"fmt"
)

func main() {
	// A map with the string: string type looks a lot
	// like a JSON object. However, note the trailing
	// comma on the final value. It is required.
	myMap := map[string]string{
		"don":  "johnson",
		"doop": "moog",
	}
	// Additional values can be added by assignment
	myMap["foo"] = "bar"
	fmt.Println(myMap)
	// Values can be iterated over using for: range
	for key, val := range myMap {
		fmt.Println(key, val)
		if key == "foo" {
			fmt.Println("Aw look it's foo")
		}
	}
}

/* OUTPUT:

map[don:johnson doop:moog foo:bar]
don johnson
doop moog
foo bar
Aw look it's foo

*/
