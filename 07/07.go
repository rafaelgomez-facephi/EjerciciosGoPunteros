package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Model string `json:"modelo"`
	Brand string `json:"marca"`
}

func printStruct(car1 *Car) {
	result, err := json.Marshal(car1)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println("keywords after---> ", string(result))
}

func main() {

	car1 := &Car{
		Model: "ibiza",
		Brand: "seat",
	}
	printStruct(car1)
}
