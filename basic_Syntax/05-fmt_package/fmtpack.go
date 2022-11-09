package main

import (
	"fmt"
)

var name string = "alex"

func main() {
	fmt.Println("Ingres your name: ")
	fmt.Scanln(&name)
	fmt.Println("Another name", name)
}
