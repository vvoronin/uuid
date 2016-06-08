package uuid


import (
	"github.com/twinj/uuid/version"
)

var _ UUID = &Unmodifiable{}

// A clean UUID type for simpler UUID versions
type Unmodifiable struct {
	*array
}

func (Unmodifiable) Size() int {
	return length
}

func (o Unmodifiable) Version() version.Version {
	return o.array.Version()
}

func (o Unmodifiable) Variant() uint8 {
	return o.array.Variant()
}

func (o Unmodifiable) Bytes() (u []byte) {
	u = append(u, *o.array...)
	return
}

func (o Unmodifiable) MarshalBinary() ([]byte, error) {
	return o.Bytes(), nil
}

func (o Unmodifiable) UnmarshalBinary(pData []byte) error {
	return nil
}

func (o Unmodifiable) HashName() (name Name) {
	return o.array.HashName()
}

func (o Unmodifiable) String() string {
	return o.array.String()
}
