package uuid

/****************
 * Date: 15/02/14
 * Time: 12:49 PM
 ***************/

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	array_bytes = [length]byte{
		0xAA, 0xCF, 0xEE, 0x12,
		0xD4, 0x00,
		0x27, 0x23,
		0x00,
		0xD3,
		0x23, 0x12, 0x4A, 0x11, 0x89, 0xFF,
	}
	idString = "aacfee12-d400-2723-00d3-23124a1189ff"
)

func TestArray_Bytes(t *testing.T) {
	id := array(array_bytes)
	assert.Equal(t, array_bytes[:], id.Bytes(), "Bytes should be the same")
}

func TestArray_Format(t *testing.T) {
	id := array(array_bytes)
	assert.Equal(t, idString, id.Format(string(CleanHyphen)), "The Format given should match the output")
}

func TestArray_MarshalBinary(t *testing.T) {
	id := array(array_bytes)
	bytes, err := id.MarshalBinary()
	assert.Nil(t, err, "There should be no error")
	assert.Equal(t, array_bytes[:], bytes, "Byte should be the same")
}

func TestArray_Size(t *testing.T) {
	id := &array{}
	assert.Equal(t, 16, id.Size(), "The size of the array should be sixteen")
}

func TestArray_String(t *testing.T) {
	id := array(array_bytes)
	assert.Equal(t, idString, id.String(), "The Format given should match the output")
}

func TestArray_UnmarshalBinary(t *testing.T) {

	u := new(array)

	err := u.UnmarshalBinary([]byte{1, 2, 3, 4, 5})

	assert.Equal(t, "uuid.UnmarshalBinary: invalid length", err.Error(), "Expect length error")

	err = u.UnmarshalBinary(array_bytes[:])

	assert.Nil(t, err, "There should be no error but got %s", err)

	for k, v := range namespaces {
		id, _ := Parse(v)
		uuidId := &array{}
		uuidId.UnmarshalBinary(id.Bytes())

		assert.Equal(t, id.Bytes(), uuidId.Bytes(), "The array id should equal the uuid id")
		assert.Equal(t, k.Bytes(), uuidId.Bytes(), "The array id should equal the uuid id")
	}
}

func TestArray_Variant(t *testing.T) {
	for _, v := range namespaces {
		id, _ := Parse(v)
		uuidId := &array{}
		uuidId.UnmarshalBinary(id.Bytes())

		assert.NotEqual(t, 0, uuidId.Variant(), "The variant should be non zero")
	}
}

func TestArray_Version(t *testing.T) {
	for _, v := range namespaces {
		id, _ := Parse(v)
		uuidId := &array{}
		uuidId.UnmarshalBinary(id.Bytes())

		assert.NotEqual(t, 0, uuidId.Version(), "The version should be non zero")
	}
}



