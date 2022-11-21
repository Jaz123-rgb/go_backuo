package main

import "fmt"

func main() {
	a := 4
	b := 5

	var r bool

	//? Igual ==
	r = a == b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)

	//? Distinto !=
	r = a != b
	fmt.Printf("a: %d es diferente que b %d? resultaado = %t \n", a, b, r)

	//? Mayor que  >
	r = a > b
	fmt.Printf("a: %d es mayor que  b %d resultaado = %t \n", a, b, r)

	//? Mayor o igual que >=
	r = a >= b
	fmt.Printf("a: %d es menor o igual que  b %d resultaado = %t \n", a, b, r)

	//? Menor que <
	r = a < b
	fmt.Printf("a: %d es menor que  b %d resultaado = %t \n", a, b, r)

	//? Menor o igual que <=
	r = a <= b
	fmt.Printf("a: %d es menor o igual que  b %d resultaado = %t \n", a, b, r)

}
