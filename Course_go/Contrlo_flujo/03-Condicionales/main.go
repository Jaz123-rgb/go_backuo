package main

import (
	"fmt"
)

func main() {
	var consumo, descuento float64
	var datosDescuento string
	// Entrada de datos
	fmt.Print("Ingress consume:  ")
	fmt.Scanln(&consumo)

	//* Condicional if-else
	if consumo >= 0 {
		if consumo <= 100 {
			//Desceunto de %10
			datosDescuento = "10%"
			descuento = consumo * 0.10

		} else if consumo >= 100 && consumo <= 200 {
			//Descuento de %20
			datosDescuento = "20%"
			descuento = consumo * 0.20

		} else if consumo >= 200 && consumo <= 300 {
			//Descuento de %30
			datosDescuento = "30%"
			descuento = consumo * 0.30

		}

	} else {
		fmt.Println("Erro to ingress numbers")
	}
	//Operaciones
	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	//Salida de datos
	fmt.Println("----------Descuento------")
	fmt.Println("Descuento que aplica", datosDescuento)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descuento: ", descuento)
	fmt.Println("Monsto con descuento", montoDescuento)
	fmt.Println("IGV", igv)
	fmt.Println("Total a pagar", totalPagar)

}
