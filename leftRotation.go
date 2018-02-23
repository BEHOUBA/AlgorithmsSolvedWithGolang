package main

import (
	"fmt"
)

func lefRotation(array []int, a int) []int {
	// var arrayResult []int
	var result []int
	result = array
	for i := 0; i < a; i++ {
		result = append(result, array[i])
	}
	return result[a:]
}

func main() {
	fmt.Println(lefRotation([]int{1, 2, 3, 4, 5}, 2))
}
