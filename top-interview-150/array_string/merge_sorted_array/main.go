package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for num1 := range nums1 {
		fmt.Println(num1)
	}
}

// Hello returns a greeting for the named person.
func main() {
	// Return a greeting that embeds the name in a message.
	merge([]int{1, 2}, 2, []int{3, 4}, 2)
}
