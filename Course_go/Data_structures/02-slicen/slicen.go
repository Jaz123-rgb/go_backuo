package main

import (
	"fmt"
)

func main() {
	//? Slicen
	num := []int{1, 2, 3}
	fmt.Println(num)

	//! Agregar dato
	num = append(num, 4)
	num = append(num, 5)
	fmt.Println("Slicen despues de usar append", num)

	//* Sub Slicen
	SubSlice := num[:2]
	num[0] = 21

	fmt.Println("Slice heredado de num", SubSlice)
	fmt.Println("Slicen num Modificado ", num)

	//* Ejercicio practo slicen
	meses := []string{"Enero", "Febrero", " Marzo"}
	fmt.Printf("len: %v, cap: %v, %p \n", len(meses), cap(meses), meses)
	meses = append(meses, "newadd")
	fmt.Printf("len: %v, cap: %v, %p \n", len(meses), cap(meses), meses)

}
