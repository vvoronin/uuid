package uuid

import "testing"

func BenchmarkNewV12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewV1() // Sets up initial store on first run
	}
	b.StopTimer()
	b.ReportAllocs()
}
