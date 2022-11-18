package parenthesis_test

import (
	"katas/10_parenthesis"
	"testing"
)

func TestEmpty(t *testing.T) {
	result := parenthesis.Run("")
	if result != true {
		t.Fail()
	}
}

func TestSimple(t *testing.T) {
	result := parenthesis.Run("()")
	if result != true {
		t.Fail()
	}
}

func TestComplexCorrect(t *testing.T) {
	result := parenthesis.Run("{ [ a * ( c + d ) ] - 5 }")
	if result != true {
		t.Fail()
	}
}

func TestComplexIncorrect(t *testing.T) {
	cases := []string{
		"{a + b [c] * (2x2)}}}}",
		"{ a * ( c + d ) ] - 5 }",
		"[ { a * ( c + d ) ] - 5 }",
		"{a^4 + (((ax4)}",
		"{ ] a * ( c + d ) + ( 2 - 3 )[ - 5 }",
		"{{{{{{(}}}}}}",
		"(a",
	}

	for _, c := range cases {
		result := parenthesis.Run(c)
		if result != false {
			t.Logf(c)
			t.Fail()
		}
	}
}
