package main

import "fmt"

func cambiaValor(x *int, y *int) {
	z := *x
	*x = *y
	*y = z
}

func orderSlice(slice []int) {
	length := len(slice)
	for i := length; i > 0; i-- {
		for j := 0; j < length; j++ {
			index := j
			index++
			if index < length {
				val1 := &slice[j]
				val2 := &slice[index]
				if *val1 > *val2 {
					cambiaValor(val1, val2)
				}
			}
		}
	}
}

func main() {
	list := []int{5, 2, 9, 3, 1, 8, 6}
	fmt.Println("Before ---> ", list)
	orderSlice(list)
	fmt.Println("After ---> ", list)
}
