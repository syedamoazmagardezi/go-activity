package main

import "fmt"

func linearSearch(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{12, 7, 25, 18, 9}

	// target exists
	result1 := linearSearch(numbers, 18)
	fmt.Println("Search 18, found at:", result1)

	// target does not exist
	result2 := linearSearch(numbers, 100)
	fmt.Println("Search 100, result:", result2)
}
