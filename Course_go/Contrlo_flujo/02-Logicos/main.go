package main

import "fmt"

func main() {

	//! ONT ! niega el dato
	fmt.Println("--------Tabla not---------")
	fmt.Println("true =", !true)
	fmt.Println("false = ", !false)
	//? AND &&

	fmt.Println("--------Tabla AND---------")
	fmt.Println("false + true =", false && true)
	fmt.Println("true + false =", true && false)
	fmt.Println("false + false =", false && false)
	fmt.Println("true + true =", true && true)

	//*OR ||
	fmt.Println("--------Tabla OR---------")
	fmt.Println("false + true =", false || true)
	fmt.Println("true + false =", true || false)
	fmt.Println("false + false =", false || false)
	fmt.Println("true + true =", true || true)

}
