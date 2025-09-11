package main

import "fmt"

func main() {
	s := "Hello  World"

	fmt.Println("Solution 1:", reverseWords1(s))
	fmt.Println("Solution 2:", reverseWords2(s))
	fmt.Println("Solution 3:", reverseWords3(s))
}
