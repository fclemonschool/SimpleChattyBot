package main

import "fmt"

func callMama(userName string) string {
	// write the function code block here
	return fmt.Sprintf("%s is learning how to call functions!", userName)
}

func main() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)
	fmt.Print(callMama(userName))
	// call the function directly, or within a fmt.Println statement here.
}
