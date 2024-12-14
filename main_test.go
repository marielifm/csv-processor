package main

import (
	"testing"
)

func BenchmarkProcessFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessFile()
	}
}
