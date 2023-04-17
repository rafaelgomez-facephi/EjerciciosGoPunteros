package main

import (
	"testing"
)

func TestChangeVowels(t *testing.T) {

	var useCases = []struct {
		name     string
		input    string
		expected string
	}{
		{"One", "test numero", "t#st n#m#r#"},
		{"Two", "El mejor mercado", "#l m#j#r m#rc#d#"},
		{"Three", "la palabra magica es", "l# p#l#br# m#g#c# #s"},
		{"Four", "si solo cabe una duda, es no saber", "s# s#l# c#b# #n# d#d#, #s n# s#b#r"},
	}

	for _, element := range useCases {
		changeVowels(&element.input)

		if element.input == element.expected {
			t.Logf(`Test %s passed OK`, element.name)
		} else {
			t.Errorf("Expected: %v, got: %v", element.expected, element.input)
		}
	}
}
