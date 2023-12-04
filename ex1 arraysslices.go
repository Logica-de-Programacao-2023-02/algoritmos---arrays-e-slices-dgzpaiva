package main

import "fmt"

func main() {
	array := [3]int{2, 4, 6}
	sum := 0
	for _, value := range array {
		sum += value
	}
	fmt.Println("A soma dos valores no array é:", sum)
}
