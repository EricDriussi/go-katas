package invert_test

import (
	"katas/06_invert"
	"testing"
)

func TestHola(t *testing.T) {
	result := invert.Run("Hola")
	if result != "aloH" {
		t.Fail()
	}
}

func TestHelloWorld(t *testing.T) {
	result := invert.Run("Hello World")
	if result != "dlroW olleH" {
		t.Fail()
	}
}
