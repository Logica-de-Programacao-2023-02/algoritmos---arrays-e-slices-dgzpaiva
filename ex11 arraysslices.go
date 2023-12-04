package main

import "fmt"

func main() {
	array := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha, coluna int
	fmt.Println("Informe um índice de linha (0 ou 1):")
	fmt.Scanln(&linha)
	fmt.Println("Informe um índice de coluna (0, 1 ou 2):")
	fmt.Scanln(&coluna)

	if linha >= 0 && linha < len(array) && coluna >= 0 && coluna < len(array[0]) {
		valor := array[linha][coluna]
		fmt.Printf("O valor armazenado na posição [%d][%d] é: %d\n", linha, coluna, valor)
	} else {
		fmt.Println("Índices inválidos!")
	}
}
