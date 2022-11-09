package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	//variables
	var a, b, suma int

	fmt.Println("Ingress first num:")
	fmt.Scanln(&a)

	fmt.Println("Ingres second num:")
	fmt.Scanln(&b)

	suma = sum(a, b)

	fmt.Println("The result is:", suma)

}
