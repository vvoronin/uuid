package uuid

const (
	length = 16
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

func (o Uuid) String() string {
	if printFormat == Canonical {
		return canonicalPrint(o)
	}
	return formatPrint(o, string(printFormat))
}

// ****************************************************

func (o Uuid) Restricted() UUID {
	return uuid(o)
}

func (o Uuid) unmarshal(pData []byte) {
	copy(o, pData)
}

func (o Uuid) setVersion(pVersion int) {
	o[versionIndex] &= 0x0f
	o[versionIndex] |= uint8(pVersion << 4)
}

func (o Uuid) setVariant(pVariant uint8) {
	setVariant(&o[variantIndex], pVariant)
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
