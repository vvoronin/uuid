package uuid

/****************
 * Date: 15/02/14
 * Time: 12:26 PM
 ***************/

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var id = &uuid{
	0xaacfee12,
	0xd400,
	0x2723,
	0x00,
	0xD3,
	[]byte{0x23, 0x12, 0x4A, 0x11, 0x89, 0xFF},
	length,
}

func TestUuid_Bytes(t *testing.T) {
	assert.Equal(t, array_bytes[:], id.Bytes(), "Bytes should be the same")
}

func TestUuid_Format(t *testing.T) {
	assert.Equal(t, idString, id.Format(string(CleanHyphen)), "The Format given should match the output")
}

func TestUuid_MarshalBinary(t *testing.T) {
	bytes, err := id.MarshalBinary()
	assert.Nil(t, err, "There should be no error")
	assert.Equal(t, array_bytes[:], bytes, "Byte should be the same")
}

func TestUuid_Size(t *testing.T) {
	assert.Equal(t, 16, id.Size(), "The size of the uuid should be sixteen")

}

func TestUuid_String(t *testing.T) {
	assert.Equal(t, idString, id.String(), "The Format given should match the output")
}

func TestUuid_UnmarshalBinary(t *testing.T) {

	u := &uuid{size: 16, node: make([]byte, 6)}

	err := u.UnmarshalBinary([]byte{1, 2, 3, 4, 5})

	assert.Equal(t, "uuid.UnmarshalBinary: invalid length", err.Error(), "Expect length error")

	err = u.UnmarshalBinary(array_bytes[:])

	assert.Nil(t, err, "There should be no error but got %s", err)

	for k, v := range namespaces {
		id, _ := Parse(v)
		uuidId := &uuid{size: 16, node: make([]byte, 6)}
		uuidId.UnmarshalBinary(id.Bytes())

		assert.Equal(t, id.Bytes(), uuidId.Bytes(), "The array id should equal the uuid id")
		assert.Equal(t, k.Bytes(), uuidId.Bytes(), "The array id should equal the uuid id")
	}
}

func TestUuid_Variant(t *testing.T) {
	for _, v := range namespaces {
		id, _ := Parse(v)
		uuidId := &uuid{size: 16, node: make([]byte, 6)}
		uuidId.UnmarshalBinary(id.Bytes())

		assert.NotEqual(t, 0, uuidId.Variant(), "The variant should be non zero")
	}
}

func TestUuid_Version(t *testing.T) {
	for _, v := range namespaces {

		id, _ := Parse(v)

		uuidId := &uuid{size: 16, node: make([]byte, 6)}
		uuidId.UnmarshalBinary(id.Bytes())

		assert.NotEqual(t, 0, uuidId.Version(), "The version should be non zero")
	}
}



