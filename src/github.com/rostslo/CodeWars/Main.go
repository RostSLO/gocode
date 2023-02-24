package main

import "fmt"

func Multiple3And5(number int) int {
	var sum int = 0

	if number < 3 {
		return 0
	}

	for i := 3; i < number; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	return sum

}

func main() {

	res := Multiple3And5(10)

	fmt.Println(res)
	fmt.Println("test")

}
