package main

import (
	"fmt"
)

func main() {

	var a bool = true     //?Boolean
	var b int = 5         //? Integer
	var c float32 = 3.14  //? Floating point number
	var d string = "Hola" //? String

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var x int = 500
	var y int = -4500

	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}
