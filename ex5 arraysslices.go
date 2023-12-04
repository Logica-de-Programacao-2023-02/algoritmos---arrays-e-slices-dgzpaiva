package main

import (
	"fmt"
)

func main() {
	var tamanho int
	fmt.Println("Informe o tamanho do Slice:")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		var valor int
		fmt.Printf("Informe o valor para o índice %d: ", i)
		fmt.Scanln(&valor)
		slice[i] = valor
	}

	fmt.Println("O Slice informado é:", slice)

	sum := 0
	for _, value := range slice {
		sum += value
	}

	fmt.Println("A soma dos valores no Slice é:", sum)
}
