package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// This is the input
	input := "1.345.679"
	// Erase "." from the string
	input = strings.ReplaceAll(input, ".", "")

	// Split the numbers from the string. Makes a slice.
	nums := strings.Split(input, "")

	// Make a new slice of int.
	numSlice := make([]int, len(nums))
	for idx, num := range nums {
		// Parse every element on the string into integer.
		numSlice[idx], _ = strconv.Atoi(num)

		// Every elements will be multiplied by 10 powered by the multiplier.
		multiplier := len(nums) - idx - 1
		numSlice[idx] = numSlice[idx] * int(math.Pow10(multiplier))
		fmt.Println(numSlice[idx])
	}
}
