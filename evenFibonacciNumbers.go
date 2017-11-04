package main

import "fmt"

// function to find all fibonacci number to a defined value and return all fibonacci numbers in an int slice
func allFibonacciNumbersTo(maxValue int) []int {
	fibonacciSlice := []int{1, 2}
	i := fibonacciSlice[1]
	for i < maxValue {
		newFib := fibonacciSlice[len(fibonacciSlice)-1] + fibonacciSlice[len(fibonacciSlice)-2]
		if newFib <= maxValue {
			fibonacciSlice = append(fibonacciSlice, newFib)
		}
		i = newFib
	}
	return fibonacciSlice
}

// function that make addition of all the even elements of a slice and return that sum
func sumPairsValue(slice []int) int {
	var sum int
	for _, x := range slice {
		if x%2 == 0 {
			sum += x
		}
	}
	return sum
}

func main() {
	// functions are tested by the number 4 million
	fiboSlice := allFibonacciNumbersTo(4000000)
	evenValueSum := sumPairsValue(fiboSlice)
	fmt.Println(evenValueSum)
}
