package main

import (
	"fmt"
)

func ArrayPunterosOrdenado(arr *[]int) []*int {
	return nil
}

func main() {
	num := make([]int, 5, 5)
	num[0] = 1
	num[1] = 5
	num[2] = 3
	num[3] = 7
	num[4] = 4
	arrayPunteros := ArrayPunterosOrdenado(&num)
	fmt.Println("Slice ordenado: ")
	for i := 0; i < len(arrayPunteros); i++ {
		fmt.Println(*arrayPunteros[i])
	}
}
