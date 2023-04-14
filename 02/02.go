package main

import (
	"fmt"
)

func cambiaValor(x *int, y *int) {
	z := *x
	*x = *y
	*y = z
}

func main() {
	x := 5
	y := 10

	fmt.Println("X --> ", x, " Y --> ", y)

	fmt.Println("X --> ", &x, " Y --> ", &y)
	cambiaValor(&x, &y)
	fmt.Println("X --> ", x, " Y --> ", y)

	fmt.Println("X --> ", &x, " Y --> ", &y)

}
