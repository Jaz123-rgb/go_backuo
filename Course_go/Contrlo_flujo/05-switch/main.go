package main

import "fmt"

func main() {
	var vocal string
	fmt.Println("Ingress a letter")
	fmt.Scanln(&vocal)

	// switch {
	// case vocal == "a" || vocal == "A":
	// 	fmt.Println(vocal, " is a vocal")
	// case vocal == "e" || vocal == "E":
	// 	fmt.Println(vocal, " is a vocal")
	// case vocal == "i" || vocal == "I":
	// 	fmt.Println(vocal, " is a vocal")
	// case vocal == "o" || vocal == "O":
	// 	fmt.Println(vocal, " is a vocal")
	// case vocal == "u" || vocal == "U":
	// 	fmt.Println(vocal, " is a vocal")
	// default:
	// 	fmt.Println(vocal, "is not a letters")
	// }

	switch vocal {
	case "A", "a", "E", "e", "I", "i", "O", "o", "U", "u":
		fmt.Println(vocal, "es una vocal")
	default:
		fmt.Println(vocal, "is not a letters")
	}

}
