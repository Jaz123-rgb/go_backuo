package main

import "fmt"

func main() {
	dias := make(map[int]string) //* <--Genera un map(arreglo) vacio
	fmt.Println(dias)

	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "viernes"
	dias[6] = "sabado"
	dias[7] = "domingo"

	fmt.Println(dias)
	dias[3] = "Test"
	fmt.Println(dias)

	//! Eliminar mapa 
	delete(dias, 3)
	fmt.Println(dias, len(dias))

	//? Mapap con arrays

	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{13, 15, 16}
	estudiantes["Roel"] = []int{14, 13, 17}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["Alex"][1])

}
