package uuid

import (
	"encoding/hex"
	"errors"
	"github.com/twinj/uuid/version"
	"encoding/binary"
)

const (
	length = 16

	// 3f used by RFC4122 although 1f works for all
	variantSet = 0x3f

	// rather than using 0xc0 we use 0xe0 to retrieve the variant
	// The result is the same for all other variants
	// 0x80 and 0xa0 are used to identify RFC4122 compliance
	variantGet = 0xe0

	variantIndex = 8
	versionIndex = 6
)

var _ UUID = &array{}

// A clean UUID type for simpler UUID versions
type array []byte

func (array) Size() int {
	return length
}

func (o array) Version() version.Version {
	if o.Variant() != ReservedRFC4122 {
		return version.Unknown
	}
	return version.Version(o[versionIndex] >> 4)
}

func (o array)  version() version.Version {
	return version.Version(o[versionIndex] >> 4)
}

func (o array) Variant() uint8 {
	return variant(o[variantIndex])
}

func (o array) unmarshal(pData []byte) {
	copy(o, pData[:length])
}

func (o array) Bytes() []byte {
	return o[:length]
}

func (o array) MarshalBinary() ([]byte, error) {
	return o[:length], nil
}

func (o array) UnmarshalBinary(pData []byte) error {
	if len(pData) != length {
		return errors.New("uuid.UnmarshalBinary: invalid length")
	}
	copy(o, pData[:])
	return nil
}

func (o array) HashName() (name Name) {
	n := make ([]byte, 16)
	groups := [][]byte{
		o[:4], o[4:6], o[6:8],
	}
	for _,v := range groups {

		binary.BigEndian.
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
	}
	for _, v := range groups {
		n = append(n, v...)
	}
	name = Name(append(n, o[8:length]...))
	return
}

func (o array) String() string {
	groups := [][]byte{
		o[:4], o[4:6], o[6:8], o[8:10], o[10:length],
	}

	size := o.Size() * 2
	hyphens := -1

	var id []byte

	switch Format(generator.Fmt) {
	case Clean, Curly, Bracket:
	default:
		size += 4
		hyphens = size
	}

	id = make([]byte, size)
	var b, e int
	for _, v := range groups {
		e = b + len(v)*2
		hex.Encode(id[b:e], v)
		b = e
		if b < hyphens {
			id[b] = '-'
			b++
		}
	}

	s := string(id)

	switch Format(generator.Fmt) {
	case Curly, CurlyHyphen:
		return "{" + s + "}"
	case Bracket, BracketHyphen:
		return "(" + s + ")"
	default:
		return s
	}
}

// ****************************************************

func (o array) setVersion(pVersion int) {
	o[versionIndex] &= 0x0f
	o[versionIndex] |= uint8(pVersion << 4)
}

func (o array) setVariant(pVariant uint8) {
	setVariant(&o[variantIndex], pVariant)
}

// Set the three most significant bits (bits 0, 1 and 2) of the
// sequenceHiAndVariant equivalent in the array to ReservedRFC4122.
func (o array) setRFC4122Version(pVersion uint8) {
	o[versionIndex] &= 0x0f
	o[versionIndex] |= uint8(pVersion << 4)
	o[variantIndex] &= variantSet
	o[variantIndex] |= ReservedRFC4122
}
