package main

import (
	"fmt"
	// "time"
	"unicode/utf8"
)

func main() {
	// var currentTIme = time.Now()

	// fmt.Print(currentTIme)

	const name = "dika"

	fmt.Printf("rune test: %d\n", utf8.RuneCountInString(name))
	fmt.Printf("len test: %d", len(name))
}