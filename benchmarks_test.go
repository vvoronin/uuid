package uuid_test

import (
	"github.com/twinj/uuid"
	_ "github.com/twinj/uuid/savers"
	"testing"
	_ "time"
)

func BenchmarkNewV1(b *testing.B) {
	uuid.NewV1()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.NewV1()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNewV2(b *testing.B) {
	uuid.NewV1()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.NewV2(uuid.DomainGroup)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNewV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid.NewV3(uuid.NamespaceDNS, uuid.Name("www.example.com"))
	}
}

func BenchmarkV3(b *testing.B) {
	id := uuid.NamespaceDNS.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V3(id, "www.example.com")
	}
}

func BenchmarkNewV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid.NewV4()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkNewV5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid.NewV5(uuid.NamespaceDNS, uuid.Name("www.example.com"))
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkArray_Bytes(b *testing.B) {
	id := uuid.NewV4()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Bytes()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkUuid_Bytes(b *testing.B) {
	id := uuid.NewV2(uuid.DomainGroup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Bytes()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkArray_String(b *testing.B) {
	id := uuid.NewV4()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.String()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkUuid_String(b *testing.B) {
	id := uuid.NewV2(uuid.DomainGroup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.String()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkEqual(b *testing.B) {
	s := "f3593cff-ee92-40df-4086-87825b523f13"
	id, _ := uuid.Parse(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.Equal(id, id)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkParse(b *testing.B) {
	s := "f3593cff-ee92-40df-4086-87825b523f13"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.Parse(s)
	}
	b.StopTimer()
	b.ReportAllocs()
}
