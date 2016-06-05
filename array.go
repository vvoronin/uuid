package uuid

import (
	"encoding/hex"
)

/****************
 * Date: 1/02/14
 * Time: 10:08 AM
 ***************/

const (
	variantIndex = 8
	versionIndex = 6
)

var _ UUID = &array{}

// A clean UUID type for simpler UUID versions
type array [length]byte

func (array) Size() int {
	return length
}

func (o array) Version() int {
	return int(o[versionIndex] >> 4)
}

func (o array) Variant() uint8 {
	return variant(o[variantIndex])
}

func (o array) Unmarshal(pData []byte) {
	copy(o[:], pData[:length])
}

func (o array) Bytes() []byte {
	return o[:]
}

func (o array) MarshalBinary() ([]byte, error) {
	return o.Bytes(), nil
}

func (o *array) UnmarshalBinary(pData []byte) error {
	return UnmarshalBinary(o, pData)
}

func (o array) String() string {
	groups := [][]byte{
		o[:4], o[4:6], o[6:8], o[8:10], o[10:],
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

func (o array) setVersion(pVersion uint8) {
	o[versionIndex] &= 0x0f
	o[versionIndex] |= pVersion << 4
}

func (o array) setVariant(pVariant uint8) {
	setVariant(&o[variantIndex], pVariant)
}

// Set the three most significant bits (bits 0, 1 and 2) of the
// sequenceHiAndVariant equivalent in the array to ReservedRFC4122.
func (o array) setRFC4122Variant() {
	o[variantIndex] &= variantSet
	o[variantIndex] |= ReservedRFC4122
}
