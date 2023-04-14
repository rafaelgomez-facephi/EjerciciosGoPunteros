package main

import "fmt"

func cambiaValor(x *int, y *int) {
	z := *x
	*x = *y
	*y = z
}

func orderSlice(slice *[]int) {
	length := len(*slice)
	for i := length; i > 0; i-- {
		for j := 0; j < length; j++ {
			index := j
			index++
			if index < length {
				val1 := &(*slice)[j]
				val2 := &(*slice)[index]
				if *val1 > *val2 {
					cambiaValor(val1, val2)
				}
			}
		}
	}
}

func ArrayPunterosOrdenado(arr *[]int) (arrResult []*int) {
	for i := 0; i < len(*arr); i++ {
		arrResult = append(arrResult, &(*arr)[i])
	}
	return
}

func main() {
	list := []int{5, 2, 9, 3, 1, 8, 6, 10}
	listPointer := []*int{}
	listValues := []int{}
	fmt.Println("slice before---> ", list)
	for i := 0; i < len(list); i++ {
		listPointer = append(listPointer, &list[i])
	}
	fmt.Println("result before ---> ", listPointer)
	orderSlice(&list)
	arrResult := ArrayPunterosOrdenado(&list)

	for i := 0; i < len(arrResult); i++ {
		listValues = append(listValues, *arrResult[i])
	}
	fmt.Println("slice after ---> ", listValues)
	fmt.Println("result after ---> ", arrResult)
}
