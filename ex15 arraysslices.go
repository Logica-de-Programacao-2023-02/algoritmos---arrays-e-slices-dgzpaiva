package main

import (
	"fmt"
)

func main() {
	array := [10]float64{3.2, 6.7, 8.9, 2.1, 9.5, 4.8, 7.3, 1.6, 5.5, 10.2}
	
	var maioresQueCinco []float64

	for _, value := range array {
		if value > 5 {
			maioresQueCinco = append(maioresQueCinco, value)
		}
	}

	fmt.Println("Novo Slice contendo os elementos maiores que 5:", maioresQueCinco)
}
