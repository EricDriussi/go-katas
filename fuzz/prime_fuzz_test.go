package fuzz_test

import (
	"katas/03_prime"
	"testing"
)

func FuzzPrimes(f *testing.F) {
	randNumbers := []int{1, 2, 3, 5, 8, 7, 97, 94}
	for _, num := range randNumbers {
		f.Add(num)
	}

	f.Fuzz(func(t *testing.T, inputNum int) {
		isPrime := prime.Run(inputNum)

		for i := 2; i < inputNum; i++ {
			if inputNum%i == 0 && isPrime {
				t.Fail()
			}
		}
	})
}
