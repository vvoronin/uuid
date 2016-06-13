package uuid

import (
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, VariantRFC4122, uuidId.Variant(), "The variant should be non zero")
	}

	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	for _, v := range uuidVariants {
		for i := 0; i <= 255; i++ {
			bytes[variantIndex] = byte(i)
			id := createUuid(bytes, 4, v)
			b := id[variantIndex] >> 4
			tVariantConstraint(v, b, id, t)
			assert.Equal(t, v, id.Variant(), "%x does not resolve to %x", id.Variant(), v)
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
		assert.Equal(t, One, id.Version(), "The version should be 1")
	}

	id := make(Uuid, length)

	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	assert.Equal(t, Unknown, id.Version(), "The version should be 0")

	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i++ {
			bytes[versionIndex] = byte(i)
			copy(id, bytes)
			id.setVersion(v)
			if v > 0 && v < 6 {
				assert.Equal(t, Version(v), id.Version(), "%x does not resolve to %x", id.Version(), v)
			} else {
				assert.Equal(t, Version(v), getVersion(id), "%x does not resolve to %x", getVersion(id), v)
			}
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

func getVersion(pId Uuid) Version {
	return Version(pId[versionIndex] >> 4)
}

func createUuid(pData []byte, pVersion int, pVariant uint8) Uuid {
	o := make(Uuid, length)
	copy(o, pData)
	o.setVersion(pVersion)
	o.setVariant(pVariant)
	return o
}

