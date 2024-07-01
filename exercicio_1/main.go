package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var vector [6]int
	for i := 0; i < 6; i++ {
		vector[i] = rand.Intn(60)
	}

	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}
}
