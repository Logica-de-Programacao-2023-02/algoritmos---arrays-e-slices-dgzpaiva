package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5}

	ordenado := estaOrdenado(array)

	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array NÃO está ordenado em ordem crescente.")
	}
}

func estaOrdenado(arr []int) bool {
	tamanho := len(arr)

	if tamanho == 0 || tamanho == 1 {
		return true
	}

	for i := 1; i < tamanho; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}

	return true
}
