package main

import "fmt"

func getHalf(a int) {
	fmt.Println("The result is", a/2)
}

// halfSum prints the sum of a and b and divides the result by 2
func halfSum(a int, b int) {
	getHalf(a + b)
}

// halfSubtract prints the result after substracting b from a and divides the result by 2
func halfSubstract(a int, b int) {
	getHalf(a - b)
}
