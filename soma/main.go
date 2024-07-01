package main

import (
	"fmt"
)

func main() {
	var num1, num2, result int

	fmt.Print("Digite o primeiro numero: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo numero: ")
	fmt.Scanln(&num2)

	result = num1 + num2
	fmt.Println(" o resultado da soma é: ", result)
}
