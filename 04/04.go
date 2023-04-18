package main

import "fmt"

func findMajor(num []int) *int {
	return nil
}

func main() {
	num := make([]int, 5, 5)
	num[0] = 1
	num[1] = 5
	num[2] = 3
	num[3] = 7
	num[4] = 4
	var mayor *int = findMajor(num)
	fmt.Println("El número más alto es: ", *mayor)
}
