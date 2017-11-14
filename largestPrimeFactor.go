package main

import "fmt"

func test(num int) bool {
	for x := num - 1; x > 2; x-- {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func findMaxPrime(num int) int {
	var val int
	for x := num; 2 < x; x-- {
		if test(x) {
			val = x
			return x
		}
	}
	return val
}

func main() {
	fmt.Println(test(600851475067))
}
