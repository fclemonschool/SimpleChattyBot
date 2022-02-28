package main

import "fmt"

func makeURL() (redirectURL string) {
	// Concatenate response parts into the redirectURL variable.
	// Parts are provided below:
	var (
		protocol   = "https://"
		domain     = "hyperskill.org/"
		path       = "golang-track?"
		parameters = "gems=1500"
	)

	fmt.Print()

	redirectURL = protocol + domain + path + parameters
	return redirectURL
}
