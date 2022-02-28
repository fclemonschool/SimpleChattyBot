package main

import "fmt"

func main() {
	var score int
	fmt.Scanf("%d", &score)

	// Write the code required to validate the student's score here.
	if score >= 71 {
		fmt.Println("Passed!")
	} else {
		fmt.Println("Failed!")
	}
}
