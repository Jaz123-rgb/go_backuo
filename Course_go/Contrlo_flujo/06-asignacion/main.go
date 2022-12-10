package main

import "fmt"

func main() {
	//Operadoeres de asiganacion
	b := 0
	b += 2
	fmt.Println(b) //Resta
	b -= 2
	fmt.Println(b) //resta
	b /= 2
	fmt.Println(b) //divide
	b *= 2
	fmt.Println(b) //multiplica por si mismop
	b %= 2
	fmt.Println(b) //resto de la division

	// al asignarel un ++ a la varibale se incrementa
	a := 0
	a++
	a++
	a++
	fmt.Println(a)
}
