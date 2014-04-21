// Classes
// Go does not have traditional classes. However,
// something equivalent can be created through using
// custom types and methods for programming in an
// OO-like style
package main

import (
	"fmt"
)

type Person struct {
	name       string
	age        int
	occupation string
}

func (p *Person) describe() {
	d := "%s is %d years old and works as a %s.\n"
	fmt.Printf(d, p.name, p.age, p.occupation)
}

func main() {
	var jim = Person{"Jimbo", 55, "landscaper"}
	var mary = Person{"Marylou", 41, "teacher"}
	jim.describe()
	mary.describe()
}

/* OUTPUT

Jimbo is 55 years old and works as a landscaper.
Marylou is 41 years old and works as a teacher.

*/
