package main

import "fmt"

func main() {
	names := [...]string{"Jaz", "Roel", "Jacinto"}
	//con for
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	//con for each

	frutas := [...]string{"Piña", "StrewBerry", "tu mama"}

	for ind, mapas := range frutas {
		fmt.Println(ind, mapas)
	}
}
