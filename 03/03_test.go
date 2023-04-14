package main

import (
	"testing"
)

func TestCambiaValor(t *testing.T) {

	var useCases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"One", []int{8, 6, 2, 10, 4}, []int{2, 4, 6, 8, 10}},
		{"Two", []int{3, 5, 2, 1, 4}, []int{1, 2, 3, 4, 5}},
		{"Three", []int{10, 6, 9, 8, 7}, []int{6, 7, 8, 9, 10}},
		{"Four", []int{108, 101, 130, 120, 110}, []int{101, 108, 110, 120, 130}},
	}

	for _, element := range useCases {
		orderSlice(element.input) // <--- invocar la función creada en el fichero del ejercicio

		if compareSlice(element.input, element.expected) {
			t.Logf(`Test %s passed OK`, element.name)
		} else {
			t.Errorf("Expected: %v, got: %v", element.expected, element.input)
		}
	}
}

func compareSlice(input []int, expected []int) bool {
	for i, element := range input {
		if element != expected[i] {
			return false
		}
	}
	return true
}
