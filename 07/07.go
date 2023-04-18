package main

import (
	"fmt"
)

func serializarJSON(p *Persona) []byte {
}

func main() {
	p := Persona{Nombre: "Juan", Apellido: "Pérez", Edad: 21}
	jsonBytes := serializarJSON(&p)
	if jsonBytes != nil {
		fmt.Println(string(jsonBytes))
	}
}
