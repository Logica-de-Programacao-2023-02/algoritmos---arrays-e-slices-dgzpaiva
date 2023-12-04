package main

import (
	"fmt"
)

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	var pares []int

	for _, value := range array {
		if value%2 == 0 {
			pares = append(pares, value)
		}
	}

	fmt.Println("Novo Slice contendo os elementos pares:", pares)
}
