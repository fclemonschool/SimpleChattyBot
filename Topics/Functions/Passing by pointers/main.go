package main

import "fmt"

func updateRapper(fName *string, lName *string) {
	*fName = "Kanye" // do not change the name!
	*lName = "West"  // do not change the last name!

	fmt.Sprint()
	//fmt.Println("My favorite rapper was: Kendrick Lamar")
	//fmt.Printf("My new favorite rapper is: %v %v", &fName, &lName)
}
