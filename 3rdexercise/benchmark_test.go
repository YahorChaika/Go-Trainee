package main

import (
	"testing"
)

var testString = "Hello  World"

func BenchmarkReverseWords1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseWords1(testString)
	}
}

func BenchmarkReverseWords2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseWords2(testString)
	}
}

func BenchmarkReverseWords3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseWords3(testString)
	}
}
