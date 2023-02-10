package main

import (
	"testing"
)

func BenchmarkFmtString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FmtString()
	}
}

func BenchmarkStrconvString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		StrconvString()
	}
}
