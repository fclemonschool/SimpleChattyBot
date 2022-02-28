package main

import "fmt"

func main() {
	var emoji = "❓"
	fmt.Scanf("%s", &emoji)

	// Please do not delete the emojis after the case statement, just fix the code errors.
	// Also please do not delete or change the text within the fmt.Println functions!
	switch emoji {
	case "⭕":
		fmt.Println("You haved picked the circle. Not the easiest shape!")
	case "🔺":
		fmt.Println("You haved picked the triangle. The easiest shape!")
	case "⭐":
		fmt.Println("You haved picked the star. Easier than circle, harder than triangle.")
	case "☂️":
		fmt.Println("You haved picked the umbrella. This is the hardest shape! GOOD LUCK.")
	default:
		fmt.Println("You haved picked an invalid emoji. Please try again or be eliminated from the game.")
	}
}
