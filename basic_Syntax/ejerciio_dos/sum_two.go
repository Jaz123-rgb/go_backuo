package main

import "fmt"

func div(a, b int) int {
	return a / b
}

func resto(a, b int) int {

	return a % b

}

func main() {
	var a, b, c, d int

	fmt.Println("First num:")
	fmt.Scanln(&a)

	fmt.Println("Second num:")
	fmt.Scanln(&b)

	c = div(a, b)

	d = resto(a, b)

	fmt.Println("Result of div", c)
	fmt.Println("Result of div", d)

}
