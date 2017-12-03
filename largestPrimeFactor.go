package main

import "fmt"

func isPrimeNumber(num int) bool {
	divP := num / 2
	for x := 3; x <= divP; x++ {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func findMaxPrime(num int) int {
	var val int
	div := num / 2
	for x := div; 2 < x; x-- {
		if num%x == 0 {
			if isPrimeNumber(x) {
				val = x
				return val
			}
		}
	}
	return val
}

func main() {
	fmt.Println(findMaxPrime(600851475067))
}
