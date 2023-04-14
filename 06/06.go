package main

import (
	"fmt"
	"strings"
)

func changeVowels(keywords *string) {
	vowels := "aeiou"
	// for i, v := range *keywords {
	// 	if strings.Contains(vowels, string(v)) {
	// 		(*keywords)[i] = uint8(rune(35))
	// 		fmt.Printf("%T, %T \n", (*keywords)[i], uint8(rune(35)))
	// 	}
	// }

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
