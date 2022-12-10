package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		izq := line[:len(line)/2]
		der := line[len(line)/2:]
		var repetido rune
	OUTER:
		for _, charIzq := range izq {
			for _, chaDer := range der {
				if chaDer == charIzq {
					repetido = chaDer
					break OUTER
				}
			}
		}
		fmt.Println(int(repetido))
	}
}
