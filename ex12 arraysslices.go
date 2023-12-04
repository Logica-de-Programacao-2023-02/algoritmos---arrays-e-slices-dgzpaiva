package main

import "fmt"

func main() {
	array := [5]int{3, 6, 9, 12, 15}
	
	var multiplosDeTres []int

	for _, value := range array {
		if value%3 == 0 {
			multiplosDeTres = append(multiplosDeTres, value)
		}
	}

	fmt.Println("Novo Slice contendo os múltiplos de 3:", multiplosDeTres)
}
