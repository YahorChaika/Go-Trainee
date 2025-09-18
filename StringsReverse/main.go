package main

import (
	"fmt"
	"strings"
)

// To run this application use the command go run main.go
func reverseRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseBuilder(s string) string {
	runes := []rune(s)
	var builder strings.Builder
	builder.Grow(len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	return builder.String()
}

func main() {
	fmt.Println(reverseRunes("Привет"))
	fmt.Println(reverseBuilder("¡Hola!"))
}
