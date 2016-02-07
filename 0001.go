package main

import "fmt"

func main() {
	// Subtract multiples of 15 as they are counted twice
	sum := sumMultiple(3, 999) + sumMultiple(5, 999) - sumMultiple(15, 999)
	fmt.Println(sum)
}

func sumMultiple(multiple, max int) int {
	// Use formula for sum of arithmetic sequence
	return (((max / multiple) * (multiple + (max / multiple * multiple))) / 2)
}
