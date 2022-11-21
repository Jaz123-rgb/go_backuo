package main

import "fmt"

func calcularigv(a, b float64) float64 {
	return a * b
}

func venta(a, b float64) float64 {
	return a + b
}

func main() {
	var a, b, c float64

	fmt.Println("Ingress sell:")
	fmt.Scanln(&a)

	b = calcularigv(a, 0.18)

	c = venta(a, b)

	// salida del output

	fmt.Println("----------------")
	fmt.Println("valor de venta", a)
	fmt.Println("IGV", b)
	fmt.Println("Precio de Venta", c)
	fmt.Println("----------------")

}
