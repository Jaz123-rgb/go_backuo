package main

import (
	"fmt"
)

func greetings(name string) {

	fmt.Println("Hi user ", name)
}

func sum(a, b int) int {
	return a + b

}

func main() {
	greetings("bitch's")
	r := sum(10, 20)
	fmt.Println(r)
}
