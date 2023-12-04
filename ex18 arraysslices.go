package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("Digite um número inteiro positivo:")
	fmt.Scanln(&n)

	primeArray := make([]int, 0)

	num := 2
	for len(primeArray) < n {
		if isPrime(num) {
			primeArray = append(primeArray, num)
		}
		num++
	}

	fmt.Println("Os", n, "primeiros números primos são:", primeArray)
}
