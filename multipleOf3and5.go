package main

import "fmt"

//This function take a number of type int and return all the multiples of 3 and 5
func multipleOfFiveAndTree(numb int) []int {
	slice := make([]int, 0, 10)
	for x := 0; x < numb; x++ {
		if x%3 == 0 {
			slice = append(slice, x)
		} else if x%5 == 0 {
			slice = append(slice, x)
		}
	}
	return slice
}

// This function take a slice of interger and return the sum of all the numbers
func addNumbers(numbers ...int) int {
	var sum int
	for x := range numbers {
		sum += numbers[x]
	}
	return sum
}

func main() {
	sliceOfValues := multipleOfFiveAndTree(1000)
	fmt.Println(sliceOfValues)
	fmt.Println(addNumbers(sliceOfValues...))

}
