package main

import (
	"fmt"
)

func main() {
	array := [6]float64{3.14, 2.718, 1.414, 1.732, 0.577, 2.236}

	var numero float64
	fmt.Println("Informe um número para adicionar em todas as posições do Array:")
	fmt.Scanln(&numero)

	for i := range array {
		array[i] += numero
	}

	fmt.Println("Array resultante:", array)
}
