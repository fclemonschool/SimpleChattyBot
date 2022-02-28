package main

import "fmt"

func main() {
	// put your code here
	var input int
	fmt.Scan(&input)

	for input > 0 {
		fmt.Print(input % 10)
		input /= 10
	}

}
