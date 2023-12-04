package main

import (
	"fmt"
)

func main() {
	array := [7]int{}

	var numero int
	fmt.Println("Informe um número para adicionar ao primeiro e ao último elemento do Array:")
	fmt.Scanln(&numero)

	array[0] += numero
	array[len(array)-1] += numero

	fmt.Println("Array resultante:", array)
}
