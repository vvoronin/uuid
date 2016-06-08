package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/twinj/uuid/version"
)

func TestArray_Bytes(t *testing.T) {
	id := make(array, length)
	copy(id, namespaceDNS.Bytes())
	assert.Equal(t, id.Bytes(), namespaceDNS.Bytes(), "Bytes should be the same")
}

func TestArray_Unmarshal(t *testing.T) {
	id := array(uuidBytes)
	id2 := make(array, length)
	id2.unmarshal(uuidBytes)

	assert.Equal(t, id.String(), id2.String(), "String should be the same")
}

func TestArray_MarshalBinary(t *testing.T) {
	id := array(uuidBytes)
	bytes, err := id.MarshalBinary()
	assert.Nil(t, err, "There should be no error")
	assert.Equal(t, uuidBytes[:], bytes, "Byte should be the same")
}

func TestArray_Size(t *testing.T) {
	id := &array{}
	assert.Equal(t, 16, id.Size(), "The size of the array should be sixteen")
}

func TestArray_String(t *testing.T) {
	id := array(uuidBytes)
	assert.Equal(t, idString, id.String(), "The Format given should match the output")
}

func TestArray_UnmarshalBinary(t *testing.T) {

	u := make(array, length)

	err := u.UnmarshalBinary([]byte{1, 2, 3, 4, 5})

	assert.Equal(t, "uuid.UnmarshalBinary: invalid length", err.Error(), "Expect length error")

	err = u.UnmarshalBinary(uuidBytes[:])

	assert.Nil(t, err, "There should be no error but got %s", err)

	for k, v := range namespaces {
		id, _ := Parse(v)
		uuidId := make(array, length)
		uuidId.UnmarshalBinary(id.Bytes())

		assert.Equal(t, id.Bytes(), uuidId.Bytes(), "The array id should equal the uuid id")
		assert.Equal(t, k.Bytes(), uuidId.Bytes(), "The array id should equal the uuid id")
	}
}

func TestArray_Variant(t *testing.T) {
	for _, v := range namespaces {
		id, _ := Parse(v)
		uuidId := make(array, length)
		uuidId.UnmarshalBinary(id.Bytes())

		assert.NotEqual(t, 0, uuidId.Variant(), "The variant should be non zero")
	}

	bytes := make(array, length)
	copy(bytes, uuidBytes[:])

	for _, v := range uuidVariants {
		for i := 0; i <= 255; i++ {
			bytes[variantIndex] = byte(i)
			id := createArray(bytes, 4, v)
			b := id[variantIndex] >> 4
			tVariantConstraint(v, b, id, t)
			output(id)
			assert.Equal(t, v, id.Variant(), "%x does not resolve to %x", id.Variant(), v)
			output("\n")
		}
	}

	assert.True(t, didArraySetVariantPanic(bytes[:]), "Array creation should panic  if invalid variant")
}

func didArraySetVariantPanic(bytes []byte) bool {
	return func() (didPanic bool) {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()

		createArray(bytes[:], 4, 0xbb)
		return
	}()
}

func TestArray_Version(t *testing.T) {
	for k, _ := range namespaces {
		id := make(array, length)
		id.UnmarshalBinary(k.Bytes())
		assert.Equal(t, version.One, id.Version(), "The version should be 1")
	}

	id := make(array, length)

	bytes := make(array, length)
	copy(bytes, uuidBytes[:])

	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i++ {
			bytes[versionIndex] = byte(i)
			id.unmarshal(bytes)
			id.setVersion(v)
			output(id)
			assert.Equal(t, version.Version(v), id.version(), "%x does not resolve to %x", id.Version(), v)
			output("\n")
		}
	}
}

// *******************************************************

func createArray(pData []byte, pVersion int, pVariant uint8) array {
	o := make(array, length)
	o.unmarshal(pData)
	o.setVersion(pVersion)
	o.setVariant(pVariant)
	return o
}
