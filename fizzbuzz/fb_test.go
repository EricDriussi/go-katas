package fizzbuzz_test

import (
	"katas/fizzbuzz"
	"testing"
)

func TestRange(t *testing.T) {
	result := fizzbuzz.Run()
	if result[0] != 1 || len(result) != 100 {
		t.Fail()
	}
}

func TestFizz(t *testing.T) {
	result := fizzbuzz.Run()
	if result[2] != "Fizz" {
		t.Fail()
	}
}

func TestBuzz(t *testing.T) {
	result := fizzbuzz.Run()
	if result[4] != "Buzz" {
		t.Fail()
	}
}

func TestFizzBuzz(t *testing.T) {
	result := fizzbuzz.Run()
	if result[14] != "FizzBuzz" {
		t.Fail()
	}
}
