package uuid

import (
	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid/version"
	"testing"
)

func TestUuid_Bytes(t *testing.T) {
	id := make(Uuid, length)
	copy(id, NameSpaceDNS.Bytes())
	assert.Equal(t, id.Bytes(), NameSpaceDNS.Bytes(), "Bytes should be the same")
}

func TestUuid_Size(t *testing.T) {
	id := make(Uuid, length)
	assert.Equal(t, 16, id.Size(), "The size of the array should be sixteen")
}

func TestUuid_String(t *testing.T) {
	id := Uuid(uuidBytes)
	assert.Equal(t, idString, id.String(), "The Format given should match the output")
}

func TestUuid_Variant(t *testing.T) {
	for _, v := range namespaces {
		id, _ := Parse(v)
		uuidId := make(Uuid, length)
		uuidId.unmarshal(id.Bytes())

		assert.NotEqual(t, 0, uuidId.Variant(), "The variant should be non zero")
	}

	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	for _, v := range uuidVariants {
		for i := 0; i <= 255; i++ {
			bytes[variantIndex] = byte(i)
			id := createUuid(bytes, 4, v)
			b := id[variantIndex] >> 4
			tVariantConstraint(v, b, id, t)
			output(id)
			assert.Equal(t, v, id.Variant(), "%x does not resolve to %x", id.Variant(), v)
			output("\n")
		}
	}

	assert.True(t, didUuidSetVariantPanic(bytes[:]), "Array creation should panic  if invalid variant")
}

func didUuidSetVariantPanic(bytes []byte) bool {
	return func() (didPanic bool) {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()

		createUuid(bytes[:], 4, 0xbb)
		return
	}()
}

func TestUuid_Version(t *testing.T) {
	for k, _ := range namespaces {
		id := make(Uuid, length)
		id.unmarshal(k.Bytes())
		assert.Equal(t, version.One, id.Version(), "The version should be 1")
	}

	id := make(Uuid, length)

	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	assert.Equal(t, version.Unknown, id.Version(), "The version should be 0")

	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i++ {
			bytes[versionIndex] = byte(i)
			copy(id, bytes)
			id.setVersion(v)
			output(id)
			assert.Equal(t, version.Version(v), getVersion(id), "%x does not resolve to %x", getVersion(id), v)
			output("\n")
		}
	}
}

func TestUuid_Restricted(t *testing.T) {
	id := NewV1()
	bb := id.Bytes()

	assert.NotNil(t, id)

	rr := id.Restricted()

	assert.Equal(t, bb, rr.Bytes(), "Bytes should be the same")

	v, ok := rr.(NameSpace)
	assert.Equal(t, v, NameSpace(""), "Should be default value")
	assert.False(t, ok, "Should not be a namespace")

	v2, ok := rr.(Uuid)
	assert.Equal(t, v2, Uuid(nil), "Should be default value")
	assert.False(t, ok, "Should not be a Uuid")
}

func TestArray_Format(t *testing.T) {

}

// *******************************************************

func getVersion(pId Uuid) version.Version {
	return version.Version(pId[versionIndex] >> 4)
}

func createArray(pData []byte, pVersion int, pVariant uint8) array {
	o := array{}
	copy(o[:], pData)
	Uuid(o[:]).setVersion(pVersion)
	Uuid(o[:]).setVariant(pVariant)
	return o
}

func createUuid(pData []byte, pVersion int, pVariant uint8) Uuid {
	o := make(Uuid, length)
	copy(o, pData)
	o.setVersion(pVersion)
	o.setVariant(pVariant)
	return o
}

// *******************************************************

func BenchmarkUuid_Bytes(b *testing.B) {
	id := make(Uuid, length)
	id.unmarshal(NameSpaceDNS.Bytes())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Bytes()
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkUuid_String(b *testing.B) {
	id := NewV2(DomainGroup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = id.String()
	}
	b.StopTimer()
	b.ReportAllocs()
}
