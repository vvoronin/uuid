package uuid

import "fmt"

const (
	length       = 16
	variantIndex = 8
	versionIndex = 6
)

// **************************************************** Default implementation

var _ UUID = &Uuid{}

type Uuid []byte

func (o Uuid) Size() int {
	return len(o)
}

func (o Uuid) Version() Version {
	return resolveVersion(o[versionIndex] >> 4)
}

func (o Uuid) Variant() uint8 {
	return variant(o[variantIndex])
}

func (o Uuid) Bytes() []byte {
	return o[:o.Size()]
}

// String returns the canonical string representation of the UUID or the
// uuid.Format the package is sent to
func (o Uuid) String() string {
	if printFormat == Canonical {
		return canonicalPrint(o)
	}
	return formatPrint(o, string(printFormat))
}

// ****************************************************

func (o Uuid) setVersion(pVersion int) {
	o[versionIndex] &= 0x0f
	o[versionIndex] |= uint8(pVersion << 4)
}

func (o Uuid) setVariant(pVariant uint8) {
	setVariant(&o[variantIndex], pVariant)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface
func (o Uuid) MarshalBinary() ([]byte, error) {
	return o.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It will return error if the slice isn't 16 bytes long.
func (o Uuid) UnmarshalBinary(pBytes []byte) (err error) {
	if len(pBytes) != o.Size() {
		err = fmt.Errorf("uuid.Uuid.UnmarshalBinary: length of bytes given [%d] must match length of Uuid going to", len(pBytes))
		return
	}
	copy(o, pBytes)

	return
}

// **************************************************** Create UUIDs

type array [length]byte

func (o *array) unmarshal(pData []byte) {
	copy(o[:], pData)
}

// Set the three most significant bits (bits 0, 1 and 2) of the
// sequenceHiAndVariant equivalent in the array to ReservedRFC4122.
func (o *array) setRFC4122Version(pVersion uint8) {
	o[versionIndex] &= 0x0f
	o[versionIndex] |= uint8(pVersion << 4)
	o[variantIndex] &= variantSet
	o[variantIndex] |= VariantRFC4122
}

// **************************************************** Immutable UUID

var _ UUID = new(Immutable)

type Immutable string

func (o Immutable) Size() int {
	return len(o)
}

func (o Immutable) Version() Version {
	return resolveVersion(o[versionIndex] >> 4)
}

func (o Immutable) Variant() uint8 {
	return variant(o[variantIndex])
}

func (o Immutable) Bytes() []byte {
	return Uuid(o).Bytes()
}

func (o Immutable) String() string {
	return Uuid(o).String()
}
