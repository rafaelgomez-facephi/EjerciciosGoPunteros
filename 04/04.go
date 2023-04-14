package main

import "fmt"

func findMajor(slice []int) (number *int) {
	number = &slice[0]
	for i, _ := range slice {
		if *number < slice[i] {
			number = &slice[i]
		}
	}
	return
}

func main() {
	list := []int{5, 2, 9, 3, 1, 8, 6, 10}
	fmt.Println("memoria before---> ", &list[7])
	number := findMajor(list)
	fmt.Println("memoria after ---> ", number)
}
