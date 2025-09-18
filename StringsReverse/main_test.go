package main

import "testing"

func BenchmarkReverseRunes(b *testing.B) {
	s := "Hola"
	for i := 0; i < b.N; i++ {
		_ = reverseRunes(s)
	}
}

func BenchmarkReverseBuilder(b *testing.B) {
	s := "Hola"
	for i := 0; i < b.N; i++ {
		_ = reverseBuilder(s)
	}
}
