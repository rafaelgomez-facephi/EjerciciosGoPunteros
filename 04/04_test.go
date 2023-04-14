package main

import (
	"testing"
)

func TestFindMajor(t *testing.T) {

	var useCases = []struct {
		name  string
		input []int
	}{
		{"One", []int{11, 6, 2, 10, 4}},
		{"Two", []int{6, 5, 2, 1, 4}},
		{"Three", []int{10, 6, 9, 8, 7}},
		{"Four", []int{140, 101, 130, 120, 110}},
	}

	for i, element := range useCases {
		memory := findMajor(element.input)

		if memory == &element.input[0] {
			t.Logf(`Test %s passed OK`, element.name)
		} else {
			t.Errorf("Expected: %v, got: %v", &element.input[i], memory)
		}
	}
}
