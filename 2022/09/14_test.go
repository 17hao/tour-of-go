package main

import "testing"

func BenchmarkWriteToMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteToMap()
	}
}

func BenchmarkWriteToSyncMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteToSyncMap()
	}
}
