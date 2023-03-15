package main

import "fmt"

func suma(x int, y int, resultado *int) {
    *resultado = x + y
}

func suma1(x int, y int) *int {
    var resultado * int = new(int)
	*resultado = x + y
	return resultado
}

func main() {
	var resultado * int = new(int)
	suma(3,6,resultado)
    fmt.Println("Pimera suma: ", *resultado)
	var resultado1 * int
	resultado1 = suma1(3,6)
	fmt.Println("Segunda suma: ", *resultado1)
}