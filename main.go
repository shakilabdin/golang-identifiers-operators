package main

import "fmt"

var y = 20 + 16

// declare the value of variable to be interger/0
var z int

func main() {
	// := short decleration operator
	// declare a variable and assign a value
	x := 8
	fmt.Println(x)

	x = 24
	fmt.Println(x)

	// var can be used to declare a variable outside the func ln 5
	// var y = 20 + 16
	fmt.Println(y)

	// declared on line 8
	z = 199 * 45
	fmt.Println(z)

}
