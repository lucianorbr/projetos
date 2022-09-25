package main

import (
	"fmt"
)

func main() {
	var alcool, gasolina float64

	fmt.Println("Imporme o valor da Gasolina: ")
	fmt.Scan(&gasolina)

	fmt.Println("Imporme o valor da Alcool: ")
	fmt.Scan(&alcool)

	if alcool/gasolina >= 0.7 {
		fmt.Println("Abastecer com Gasolina")
	} else {
		fmt.Println("Abastecer com Álcool")
	}

}
