package fizzbuzz

import "fmt"

func Run() []interface{} {
	var fbarray []interface{}

	for i := 1; i <= 100; i++ {
		divByThree := i%3 == 0
		divByFive := i%5 == 0
		if divByFive && divByThree {
			fbarray = append(fbarray, "FizzBuzz")
		} else if divByThree {
			fbarray = append(fbarray, "Fizz")
		} else if divByFive {
			fbarray = append(fbarray, "Buzz")
		} else {
			fbarray = append(fbarray, i)
		}
	}
	printArr(fbarray)
	return fbarray
}

func printArr(arr []interface{}) {
	for _, e := range arr {
		fmt.Println(e)
	}
	fmt.Println()
}
