package uuid

import (
	"testing"
	_ "time"
)

var name UniqueName = Name("www.widgets.com")

func init() {
	Init()
}

func BenchmarkNewV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewV1() // Sets up initial store on first run
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNewV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewV2(DomainGroup)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNewV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewV3(NameSpaceDNS, name)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNewV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewV4()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNewV5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewV5(NameSpaceDNS, name)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNameSpace_Bytes(b *testing.B) {
	id := make(Uuid, length)
	id.unmarshal(NameSpaceDNS.Bytes())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Bytes()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkEqual(b *testing.B) {
	s := "f3593cff-ee92-40df-4086-87825b523f13"
	id, _ := Parse(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Equal(id, id)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkParse(b *testing.B) {
	s := "f3593cff-ee92-40df-4086-87825b523f13"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(s)
	}
	b.StopTimer()
	b.ReportAllocs()
}
