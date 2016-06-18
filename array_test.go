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

func TestImmutable_Bytes(t *testing.T) {
	b := make([]byte, length)
	copy(b[:], NameSpaceDNS.Bytes())

	id := Immutable(b)

	assert.Equal(t, NameSpaceDNS.Bytes(), id.Bytes())
}

func TestImmutable_Size(t *testing.T) {
	assert.Equal(t, 16, Nil.Size(), "The size of the array should be sixteen")
}

func TestImmutable_String(t *testing.T) {
	id := Immutable(uuidBytes)
	assert.Equal(t, idString, id.String(), "The Format given should match the output")
}

func TestImmutable_Variant(t *testing.T) {
	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	for _, v := range uuidVariants {
		for i := 0; i <= 255; i++ {
			bytes[variantIndex] = byte(i)
			id := createUuid(bytes, 4, v)
			b := id[variantIndex] >> 4
			tVariantConstraint(v, b, id, t)
			id2 := Immutable(id)
			assert.Equal(t, v, id2.Variant(), "%x does not resolve to %x", id2.Variant(), v)
		}
	}
}

func TestImmutable_Version(t *testing.T) {

	id := make(Uuid, length)
	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i++ {
			bytes[versionIndex] = byte(i)
			copy(id, bytes)
			id.setVersion(v)
			id2 := Immutable(id)

			if v > 0 && v < 6 {
				assert.Equal(t, Version(v), id2.Version(), "%x does not resolve to %x", id2.Version(), v)
			} else {
				assert.Equal(t, Version(v), getVersion(Uuid(id)), "%x does not resolve to %x", getVersion(Uuid(id)), v)
			}
		}
	}
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
