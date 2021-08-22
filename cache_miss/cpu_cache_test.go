package cache_miss

import "testing"

func BenchmarkPrintRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintRow()
	}
}

func BenchmarkPrintColumn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintColumn()
	}
}
