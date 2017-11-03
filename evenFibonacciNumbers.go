package main

import "fmt"

func allFibonacciNumbersTo(num int) []int {
	slice := make([]int, 1, 10)
	slice[0] = 1
	slice[1] = 2
	x := 0
	for x < num {
		y := slice[-1] + slice[-2]
		slice = append(slice, y)
		x = slice[len(slice)-1]
	}
	return slice
}

func main() {

	fmt.Println(allFibonacciNumbersTo(40))
}
