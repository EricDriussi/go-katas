package prime

import "fmt"

func Run(num int) bool {
	printPrimeNumbersTo100()
	return isPrime(num)
}

func printPrimeNumbersTo100() {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
