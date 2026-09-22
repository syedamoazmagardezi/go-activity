package main

import "fmt"

func binarySearch(numbers []int, target int) int {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		middle := (left + right) / 2
		fmt.Println("Checking index", middle, "value", numbers[middle])

		if numbers[middle] == target {
			return middle
		} else if numbers[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func main() {
	numbers := []int{3, 7, 9, 12, 18, 25, 31}

	fmt.Println("Result for 18:", binarySearch(numbers, 18))
	fmt.Println("---")
	fmt.Println("Result for 3:", binarySearch(numbers, 3))
	fmt.Println("---")
	fmt.Println("Result for 31:", binarySearch(numbers, 31))
	fmt.Println("---")
	fmt.Println("Result for 100:", binarySearch(numbers, 100))
}
