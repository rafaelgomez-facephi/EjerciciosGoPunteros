package main

import (
	"testing"
)

func TestCambiaValorOK(t *testing.T) {

	var useCases = []struct {
		name string
		a    int
		b    int
	}{
		{"One", 4, 2},
		{"Two", 10, 5},
		{"Three", 20, 10},
		{"Four", 30, 15},
	}

	for _, element := range useCases {
		expectedA := element.b
		expectedB := element.a

		cambiaValor(&element.a, &element.b) // <--- invocar la función creada en el fichero del ejercicio

		if expectedA == element.a && expectedB == element.b {
			t.Logf(`Test %s passed OK`, element.name)
		} else {
			t.Fail()
			t.Errorf("Expected: %v, got: %v", expectedA, element.a)
		}
	}
}

func TestCambiaValorKO(t *testing.T) {

	var useCases = []struct {
		name string
		a    int
		b    int
	}{
		{"One", 4, 2},
		{"Two", 10, 5},
		{"Three", 20, 10},
		{"Four", 30, 15},
	}

	for _, element := range useCases {
		expectedA := element.a
		expectedB := element.b

		cambiaValor(&element.a, &element.b)

		if expectedA != element.a && expectedB != element.b {
			t.Logf(`Test %s passed OK`, element.name)
		} else {
			t.Errorf("Expected: %v, got: %v", expectedA, element.a)
			t.Fail()
		}
	}
}
