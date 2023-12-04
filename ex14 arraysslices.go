package main

import "fmt"

func main() {
	slice := make([]int, 8)

	fmt.Println("Informe dois índices para trocar de posição no Slice:")
	var indice1, indice2 int
	fmt.Scanln(&indice1)
	fmt.Scanln(&indice2)

	if indice1 >= 0 && indice1 < len(slice) && indice2 >= 0 && indice2 < len(slice) {
		slice[indice1], slice[indice2] = slice[indice2], slice[indice1]
		fmt.Println("Slice resultante após a troca:", slice)
	} else {
		fmt.Println("Índices inválidos!")
	}
}
