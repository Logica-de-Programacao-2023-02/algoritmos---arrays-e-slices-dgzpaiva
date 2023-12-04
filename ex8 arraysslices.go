package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 8)

	slice[0] = "maçã"
	slice[1] = "banana"
	slice[2] = "laranja"
	slice[3] = "maçã"
	slice[4] = "uva"
	slice[5] = "abacaxi"
	slice[6] = "maçã"
	slice[7] = "pera"

	var valor string
	fmt.Println("Informe um valor para remover do Slice:")
	fmt.Scanln(&valor)

	resultado := make([]string, 0)

	for _, v := range slice {
		if v != valor {
			resultado = append(resultado, v)
		}
	}

	fmt.Println("Resultado após remover as ocorrências:", resultado)
}
