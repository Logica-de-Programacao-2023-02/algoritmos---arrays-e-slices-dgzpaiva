package main

import "fmt"

func main() {
	array := [4]float64{2.5, 3.5, 4.5, 5.5}
	product := 1.0

	for _, value := range array {
		product *= value
	}

	fmt.Println("O produto dos valores no array é:", product)
}
