package main

import "fmt"

func main() {
	// put your code here
	var input int
	fmt.Scan(&input)

	var result = 1
	if input == 0 {
		fmt.Print(1)
	}

	for input >= 0 {
		result *= input
		input--
		if input == 0 {
			fmt.Println(result)
		}
	}
}
