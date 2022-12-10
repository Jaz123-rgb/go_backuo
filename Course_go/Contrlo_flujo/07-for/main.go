package main

import "fmt"

func main() {

	//! En go no existe el whule
	numeros := 12455

	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("Cantidad de digios", c)

	//Bucle for
	for i := 0; i <= 100; i++ {
		if i == 2 {
			fmt.Println("soy una multiplo de 2", i)
			continue
		}
		if i == 50 {
			fmt.Println("hasta aqui llego")
			break
		}
		fmt.Println(i + 1)
	}
}
