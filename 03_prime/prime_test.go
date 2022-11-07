package prime_test

import (
	"katas/03_prime"
	"testing"
)

func TestOneNotPrime(t *testing.T) {
	result := prime.Run(1)
	if result != false {
		t.Fail()
	}
}

func TestTwoIsPrime(t *testing.T) {
	result := prime.Run(2)
	if result != true {
		t.Fail()
	}
}

func TestThreeIsPrime(t *testing.T) {
	result := prime.Run(3)
	if result != true {
		t.Fail()
	}
}

func TestSevenIsPrime(t *testing.T) {
	result := prime.Run(7)
	if result != true {
		t.Fail()
	}
}

func TestNinetySevenIsPrime(t *testing.T) {
	result := prime.Run(97)
	if result != true {
		t.Fail()
	}
}

func TestSixNotPrime(t *testing.T) {
	result := prime.Run(6)
	if result != false {
		t.Fail()
	}
}

func TestTwentyNotPrime(t *testing.T) {
	result := prime.Run(20)
	if result != false {
		t.Fail()
	}
}
