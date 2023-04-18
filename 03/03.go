package main

import "fmt"

func intercambiar(x *int, y *int) {
}

// Algoritmo de la burbuja
func orderSlice(num []int) {
}

func main() {
	num := make([]int, 5, 5)
	num[0] = 1
	num[1] = 5
	num[2] = 3
	num[3] = 7
	num[4] = 4
	orderSlice(num)
	fmt.Println("Slice ordenado: ", num)
}
