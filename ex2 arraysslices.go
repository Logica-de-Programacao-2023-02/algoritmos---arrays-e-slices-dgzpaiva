package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println("Slice resultante:", slice)
}
