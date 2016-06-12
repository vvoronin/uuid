package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameSpace_Bytes(t *testing.T) {
	b := make([]byte, length)
	copy(b[:], NameSpaceDNS.Bytes())

	NewV3(NameSpaceDNS, Name("www.widgets.com"))
	assert.Equal(t, b, NameSpaceDNS.Bytes())

	NewV3(NameSpaceDNS, Name("www.widgets.com"))
	assert.Equal(t, b, NameSpaceDNS.Bytes())

	changeOrder(NameSpaceDNS.Bytes())
	assert.Equal(t, b, NameSpaceDNS.Bytes())
}

func TestUuid_Bytes2(t *testing.T) {
	b := make([]byte, length)
	copy(b[:], NameSpaceDNS.Bytes())

	id := uuid(b)

	assert.Equal(t, NameSpaceDNS.Bytes(), id.Bytes())
}

func TestNameSpace_Size(t *testing.T) {
	assert.Equal(t, 16, NameSpaceDNS.Size(), "The size of the array should be sixteen")
}

func TestUuid_Size2(t *testing.T) {
	assert.Equal(t, 16, Nil.Size(), "The size of the array should be sixteen")
}

func TestNameSpace_String(t *testing.T) {
	id := Uuid(uuidBytes)
	id2 := PromoteToNameSpace(id)
	assert.Equal(t, idString, id2.String(), "The Format given should match the output")
}

func TestUuid_String2(t *testing.T) {
	id := uuid(uuidBytes)
	assert.Equal(t, idString, id.String(), "The Format given should match the output")
}

func TestNameSpace_Variant(t *testing.T) {
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
			id2 := PromoteToNameSpace(id)
			output(id)
			assert.Equal(t, v, id2.Variant(), "%x does not resolve to %x", id2.Variant(), v)
			output("\n")
		}
	}
}

func TestUuid_Variant2(t *testing.T) {
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
			id2 := uuid(id)
			output(id)
			assert.Equal(t, v, id2.Variant(), "%x does not resolve to %x", id2.Variant(), v)
			output("\n")
		}
	}
}

func TestNameSpace_Version(t *testing.T) {
	for k, _ := range namespaces {
		id := make(Uuid, length)
		id.unmarshal(k.Bytes())
		assert.Equal(t, One, id.Version(), "The version should be 1")
	}

	id := make(Uuid, length)

	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i++ {
			bytes[versionIndex] = byte(i)
			copy(id, bytes)
			id.setVersion(v)
			id2 := PromoteToNameSpace(id)

			output(id)
			if v > 0 && v < 6 {
				assert.Equal(t, Version(v), id2.Version(), "%x does not resolve to %x", id2.Version(), v)
			} else {
				assert.Equal(t, Version(v), getNamespaceVersion(Uuid(id2)), "%x does not resolve to %x", getNamespaceVersion(Uuid(id2)), v)
			}
			output("\n")
		}
	}
}

func TestUuid_Version2(t *testing.T) {
	for k, _ := range namespaces {
		id := make(Uuid, length)
		id.unmarshal(k.Bytes())
		assert.Equal(t, One, id.Version(), "The version should be 1")
	}

	id := make(Uuid, length)

	bytes := make(Uuid, length)
	copy(bytes, uuidBytes[:])

	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i++ {
			bytes[versionIndex] = byte(i)
			copy(id, bytes)
			id.setVersion(v)
			id2 := uuid(id)

			output(id)
			if v > 0 && v < 6 {
				assert.Equal(t, Version(v), id2.Version(), "%x does not resolve to %x", id2.Version(), v)
			} else {
				assert.Equal(t, Version(v), getVersion(Uuid(id)), "%x does not resolve to %x", getVersion(Uuid(id)), v)
			}
			output("\n")
		}
	}
}

func TestMarshaller_MarshalBinary(t *testing.T) {
	//id := Uuid(uuidBytes)
	//u := Marshaller(id)
	//bytes, err := u.MarshalBinary()
	//assert.Nil(t, err, "There should be no error")
	//assert.Equal(t, uuidBytes[:], bytes, "Byte should be the same")
}

func TestUuid_UnmarshalBinary(t *testing.T) {

	//u := make(Uuid, length)
	//
	//err := u.UnmarshalBinary([]byte{1, 2, 3, 4, 5})
	//
	//assert.Equal(t, "uuid.Marshaller.UnmarshalBinary: invalid length", err.Error(), "Expect length error")
	//
	//err = u.UnmarshalBinary(uuidBytes[:])
	//
	//assert.Nil(t, err, "There should be no error but got %s", err)
	//
	//for k, v := range namespaces {
	//	id, _ := Parse(v)
	//	u = make(Uuid, length)
	//	u.UnmarshalBinary(id.Bytes())
	//
	//	assert.Equal(t, id.Bytes(), mm.Bytes(), "The array id should equal the uuid id")
	//	assert.Equal(t, k.Bytes(), mm.Bytes(), "The array id should equal the uuid id")
	//}
	//
	//uu := uuid_array{}

	//u = Marshaller(uu[:])
	//err = u.UnmarshalBinary(uuidBytes[:])
	//assert.Nil(t, err, "There should be no error but got %s", err)
	//assert.Equal(t, uuidBytes[:], uu.Bytes(), "The array id should equal the uuid id")
	//
	//var ii UUID = V1()
	//
	//u = Marshaller(ii.(Uuid))
	//err = u.UnmarshalBinary(uuidBytes[:])
	//assert.Nil(t, err, "There should be no error but got %s", err)
	//assert.Equal(t, uuidBytes[:], ii.Bytes(), "The array id should equal the uuid id")
	//
	//jj := new(uuid)
	//
	//u = Marshaller(Uuid(*jj))
	//err = u.UnmarshalBinary(uuidBytes[:])
	//assert.Error(t, err, "There should be an error")
	//
	//u = Marshaller(Uuid(Nil))
	//err = u.UnmarshalBinary(uuidBytes[:])
	//assert.Nil(t, err, "There should be no error but got %s", err)
	//assert.NotEqual(t, uuidBytes[:], Nil.Bytes(), "The array id should equal the uuid id")
	//
	//u = Marshaller(Uuid(NameSpaceDNS))
	//err = u.UnmarshalBinary(uuidBytes[:])
	//assert.Nil(t, err, "There should be no error but got %s", err)
	//assert.NotEqual(t, uuidBytes[:], NameSpaceDNS.Bytes(), "The array id should equal the uuid id")
	//
	//kk := V1()
	//ll := kk.Unmodifiable()
	//
	//u = Marshaller(ll.(uuid))
	//err = u.UnmarshalBinary(uuidBytes[:])
	//assert.Nil(t, err, "There should be no error but got %s", err)
	//assert.Equal(t, uuidBytes[:], ll.Bytes(), "The array id should equal the uuid id")

}

func getNamespaceVersion(pId Uuid) Version {
	return Version(pId[versionIndex+1] >> 4)
}
