package main

import "fmt"

func main() {
	var age int
	fmt.Scanf("%d", &age)

	// Code your switch or if...else-if statement here.
	if age <= 14 {
		fmt.Println("Toy Story 4")
	} else if age >= 15 && age <= 18 {
		fmt.Println("The Matrix")
	} else if age >= 19 && age <= 25 {
		fmt.Println("John Wick")
	} else if age >= 26 && age <= 35 {
		fmt.Println("Constantine")
	} else if age > 35 {
		fmt.Println("Speed")
	}
}
