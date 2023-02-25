package main

import "github.com/rostslo/gocode/src/github.com/rostslo/CodeWars/calc"

func main() {

	ad := calc.Addition{1, 9}
	calc.executeOperation(&ad)

}

/*
func main() {

	res := Multiple3And5(10)

	fmt.Println(res)
	fmt.Println("test")

}

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
*/
