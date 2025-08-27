package main

import (
	"fmt"
	"unicode"
)

// ------------------------
// normalize убирает пробелы и приводит к нижнему регистру
func normalize(s string) []rune {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		out = append(out, unicode.ToLower(r))
	}
	return out
}

// ------------------------
// 1) Разворот строки
func IsPalindromeReverse(s string) bool {
	r := normalize(s)
	n := len(r)
	rev := make([]rune, n)
	for i := 0; i < n; i++ {
		rev[n-1-i] = r[i]
	}
	for i := 0; i < n; i++ {
		if r[i] != rev[i] {
			return false
		}
	}
	return true
}

// ------------------------
// 2) Два указателя
func IsPalindromeTwoPointer(s string) bool {
	r := normalize(s)
	i, j := 0, len(r)-1
	for i < j {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// ------------------------
// 3) Деление пополам
func IsPalindromeHalf(s string) bool {
	r := normalize(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		if r[i] != r[n-1-i] {
			return false
		}
	}
	return true
}

// ------------------------
// main() для проверки
func main() {
	tests := []string{
		"MrOwl ate my metal worm",
		"Never Odd Or Even",
		"1223221",
		"No lemon no melon",
		"Дом Мод",
		"Топот",
		"not a palindrome",
	}

	for _, s := range tests {
		fmt.Println(s)
		fmt.Println("Reverse:", IsPalindromeReverse(s))
		fmt.Println("TwoPointer:", IsPalindromeTwoPointer(s))
		fmt.Println("Half:", IsPalindromeHalf(s))
		fmt.Println("---")
	}
}
