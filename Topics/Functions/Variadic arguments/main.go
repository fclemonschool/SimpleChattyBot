package main

import "fmt"

func divide(nums ...float64) float64 {
	total := 1.0
	for _, num := range nums {
		total /= num
	}
	fmt.Print()
	return total
}
