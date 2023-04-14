package main

import (
	"fmt"
	"strings"
)

func changeVowels(keywords *string) {
	vowels := "aeiouAEIOU"

	for _, v := range vowels {
		*keywords = strings.ReplaceAll(*keywords, string(v), "#")
	}
}

func main() {
	keywords := "hola mundo"
	fmt.Println("keywords before---> ", keywords)
	changeVowels(&keywords)
	fmt.Println("keywords after ---> ", keywords)
}
