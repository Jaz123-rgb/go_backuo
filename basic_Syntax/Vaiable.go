package main

import (
	"fmt"
)

// var a int
// var b int = 2
// var c = 3

func main() {
	// var std1 string = "John"  type us a string
	// var std2 = "Jane"
	// x := 2

	// fmt.Println(std1)
	// fmt.Println(std2)
	// fmt.Println(x)
	// a = 1

	//? fmt.Println(a)
	//? fmt.Println(b)
	//? fmt.Println(c)

	var a, b int = 1, 2
	c, d := 7, "Word"
	var (
		ite1 int
		ite2 int    = 1
		ite3 string = "hello"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(ite1, ite2, ite3)
}

/*? difference between var an := var is global in the code
and := just inside on a function

*/

// exist difereten
