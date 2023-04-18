package main

import "fmt"

func cambiaValor(x *int, y *int) {
}

func main() {
	var x, y int = 2, 4
	cambiaValor(&x, &y)
	fmt.Printf("Ahora x = %d e y = %d\n", x, y)
}
