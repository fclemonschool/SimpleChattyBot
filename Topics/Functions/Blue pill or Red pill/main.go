package main

import "fmt"

// Do not change the code of this function!
func takePill(pill string) string {
    if pill == "red" {
        return "You stay in wonderland, and see how deep the rabbit hole goes."
    } else if pill == "blue" {
        return "The story ends, you wake up in bed and believe what you want to believe."
    } else {
        return "You wake up in a strange place, and you don't know what to do."
    }
}

func main() {
    var selection string
    fmt.Scanf("%s", &selection)

    // call the takePill function within a fmt.Println() statement here by passing the selection variable as an argument.
}