package main

import (
	"fmt"
)

func changeVowels(str *string) {
}

func main() {
	str := "Hola mundo"
	changeVowels(&str)
	fmt.Println("String sin vocales: ", str)
}
