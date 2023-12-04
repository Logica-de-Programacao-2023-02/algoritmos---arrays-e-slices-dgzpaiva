package main

import (
	"fmt"
)

func main() {
	rows, cols := 3, 2
	matrix := make([][]int, rows)

	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var value int
			fmt.Printf("Informe o valor para a posição [%d][%d]: ", i, j)
			fmt.Scanln(&value)
			matrix[i][j] = value
		}
	}

	fmt.Println("A matriz informada é:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
