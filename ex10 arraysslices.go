package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{8, 2, 10, 4, 6, 1, 7, 3, 9, 5}

	min := math.MaxInt64
	max := math.MinInt64

	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	fmt.Println("Valor mínimo no slice:", min)
	fmt.Println("Valor máximo no slice:", max)
}
