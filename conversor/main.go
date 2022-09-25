package main

import (
	"fmt"
	"os"
)

func main() {

	var temperatura float64
	var unidade string

	fmt.Println("Imporme a temperatura: ")
	fmt.Scan(&temperatura)

	fmt.Println("Informe C para Celsius e F para Fahrenheit: ")
	fmt.Scan(&unidade)

	if unidade == "C" {
		fmt.Println("A temperatura em Fahrenheit é: ", (temperatura*1.8)+32)
	} else if unidade == "F" {
		fmt.Println("A temperatura em Celsius é: ", (temperatura-32)/1.8)
	} else {
		fmt.Println("Valor inválido")
		os.Exit(1)
	}

}
