package main

func reconstructSpeech(words []string) (speech string) {
    // From the very beginning there are no words at all
    speech = ""
    for _, word := range words {
        // Each iteration you should
        speech += word // Concatenate "word" and "speech" variables
    }

    return speech
}
