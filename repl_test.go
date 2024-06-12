package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	input := "HELLO"
	expected := []string{"hello"}

	actual := cleanInput(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %v, got %v", v, actual[i])
		}
	}
}

func TestCleanInputMultipleWords(t *testing.T) {
	input := "HELLO WORLD"
	expected := []string{"hello", "world"}

	actual := cleanInput(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %v, got %v", v, actual[i])
		}
	}
}
