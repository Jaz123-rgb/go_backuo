package main

import (
	"fmt"
)

func main() {
	if name, age := "Jaz", 24; name == "Jaz" {
		fmt.Println("Hola, ", name)
	} else {
		fmt.Printf("Nombre: %s - edad %d\n", name, age)
	}
  
	user := make(map[string]string)

	user["ale"] = "alex@gmail.com"
	user["Jaz"] = "jaziel@oultook.com"

	if correo, ok := user["Jaz"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
