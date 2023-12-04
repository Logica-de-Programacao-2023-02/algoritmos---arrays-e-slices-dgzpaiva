package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5)

	var numero int
	fmt.Println("Informe um número inteiro:")
	fmt.Scanln(&numero)

	found := false
	for _, value := range slice {
		if value == numero {
			found = true
			break
		}
	}

	if !found {
		for i, value := range slice {
			if value == 0 {
				slice[i] = numero
				break
			}
		}
	}

	fmt.Println("Slice resultante:", slice)
}
