package main

import "fmt"

func main() {
	num := make([]int, 3, 3)
	num[0] = 100
	num[1] = 101
	num[2] = 102

	fmt.Println(num, len(num), cap(num))

	//* Agregar nuebvo dato al slicen
	num = append(num, 104)

	fmt.Println(num, len(num), cap(num))

}
