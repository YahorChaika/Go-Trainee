package main

import "testing"

// тестовые строки
var testStrings = []string{
	"MrOwl ate my metal worm",
	"Never Odd Or Even",
	"1223221",
	"No lemon no melon",
	"Дом Мод",
	"Топот",
	"not a palindrome",
}

// один бенчмарк с под-бенчмарками для каждого метода
func BenchmarkAll(b *testing.B) {
	for _, str := range testStrings {
		b.Run("Reverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = IsPalindromeReverse(str)
			}
		})
		b.Run("TwoPointer", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = IsPalindromeTwoPointer(str)
			}
		})
		b.Run("Half", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = IsPalindromeHalf(str)
			}
		})
	}
}
