package main

import (
	"testing"
)


func IsPalindromeReverse(s string) bool {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r) == s
}


func IsPalindromeTwoPointer(s string) bool {
	r := []rune(s)
	left, right := 0, len(r)-1
	for left < right {
		if r[left] != r[right] {
			return false
		}
		left++
		right--
	}
	return true
}


func IsPalindromeHalf(s string) bool {
	r := []rune(s)
	n := len(r)
	half := n / 2
	for i := 0; i < half; i++ {
		if r[i] != r[n-1-i] {
			return false
		}
	}
	return true
}


var testStrings = []string{
	"1223221",
	"radar",
	"level",
	"abba",
	"hello",
}


func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			IsPalindromeReverse(s)
			IsPalindromeTwoPointer(s)
			IsPalindromeHalf(s)
		}
	}
}


func main() {
	words := []string{"1223221", "radar", "abba", "hello", "level"}
	for _, w := range words {
		println(w,
			IsPalindromeReverse(w),
			IsPalindromeTwoPointer(w),
			IsPalindromeHalf(w))
	}
}
