package main

import (
	"fmt"
)

func main() {

	///array
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 12
	numeros[1] = 23
	numeros[2] = 32

	fmt.Println(numeros)
	fmt.Println(numeros[4])

	//? Array con valores

	name := [3]string{"Jaz", "Luis", "ema"}
	fmt.Println(name)

	// arrays ssin datos en especififocs
	colors := [...]string{
		"Red",
		"Green",
		"Black",
		"Blue",
	}
	fmt.Println(colors, len(colors))

	//! indices indefinidos
	coins := [...]string{
		0: "Doalr",
		1: "SOles",
		2: "Pesos",
		3: "Euro",
	}
	fmt.Println(coins, len(coins))
	//? Subarrays
	Subcoins := coins[0:2]
	fmt.Println(Subcoins)

}
