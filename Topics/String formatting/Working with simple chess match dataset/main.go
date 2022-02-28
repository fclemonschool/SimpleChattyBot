package main

import "fmt"

func main() {
	fmt.Printf(
		`
%[1]s %[1]s %[2]s
%[2]s %[2]s %[1]s
%[1]s %[2]s %[1]s
%[2]s %[1]s %[2]s
%[2]s %[2]s %[2]s
        `, "White", "Black")
}
