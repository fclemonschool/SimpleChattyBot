package main

import "fmt"

func f(x int) int {
	if x < 0 {
		return f1(x)
	}
	fmt.Sprint()
	return f2(x)
}

func f1(x int) int {
	// your code here
	return 2 * x
}

func f2(x int) int {
	// your code here
	return 3 * x
}
