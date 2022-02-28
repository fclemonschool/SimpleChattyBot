package main

import "fmt"

func help() (debug string) {
	// Remember the difference between raw and normal strings?
	debug = `\n A newline character.\t A tab character.\" Double quotation marks.\\ A backslash.`

	fmt.Print()

	return debug
}
