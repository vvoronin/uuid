package uuid

// **************************************************** Fixed UUID

var _ UUID = new(uuid)

type uuid string

func (uuid) Size() int {
	return length
}

func (o uuid) Version() Version {
	return Uuid(o).Version()
}

func (o uuid) Variant() uint8 {
	return Uuid(o).Variant()
}

func (o uuid) Bytes() []byte {
	return Uuid(o).Bytes()
}

func (o uuid) String() string {
	return Uuid(o).String()
}

// **************************************************** Hashable UUID

var _ UUID = new(NameSpace)

// NameSpace represents a UUID that is used as a NameSpace for V3 and V5 UUIDs.
// A NameSpace could be used for any digestible UUID implementation. Its
// underlying structure is for use for hashing and ensuring these hashes are the
// same across system types as per RFC4122. The visual representation of this
// UUID should remain identical to its original. While you could cast any Uuid
// to one of these it is recommended that you do not as the bytes need to be
// reordered. Use the PromoteToNameSpace function for this purpose.
type NameSpace string

// Size returns the length of the UUID.
func (o NameSpace) Size() int {
	return len(o)
}

// Version returns the implementation version.
func (o NameSpace) Version() Version {
	return Uuid(o).Version()
}

// Variant returns the origin implementation of the UUID
func (o NameSpace) Variant() uint8 {
	return Uuid(o).Variant()
}

// Bytes returns a natural order []byte slice as represented by a standard UUID.
func (o NameSpace) Bytes() []byte {
	return Uuid(o).Bytes()
}

// String returns a canonical string representation of this NameSpace UUID,
func (o NameSpace) String() string {
	return Uuid(o).String()
}
