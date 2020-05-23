package main

import "fmt"

// TwoNumberSum is to find if two numbers add up to the target number.
func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	content := make([]int, 0, 10)
	for i := 0; i < len(array); i++ {
		temp := target - array[i]
		for j := 0; j < len(array); j++ {
			if array[j] == temp && i != j {
				content = append(content, array[j])
			}
		}
	}
	return content
}

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	result := TwoNumberSum(array, targetSum)
	fmt.Println(result)
}