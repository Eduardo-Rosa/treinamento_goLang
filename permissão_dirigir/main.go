package main

import "fmt"

func main() {

	for idade := 10; idade < 18; idade++ {
		fmt.Println("Ainda não pode dirigir!")
	}
	fmt.Println("Já pode dirigir!")
}
