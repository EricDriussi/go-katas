package main

import (
	"katas/01_fizzbuzz"
	"katas/02_anagram"
	"katas/03_prime"
)

func main() {
	fizzbuzz.Run()
	anagram.Run("inch", "chin")
	prime.Run(7)
}
