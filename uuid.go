// This package provides RFC4122 and DCE 1.1 UUIDs.
//
// NewV1, NewV2, NewV3, NewV4, NewV5, for generating versions 1, 3, 4
// and 5 UUIDs as specified in RFC-4122.
//
// New([]byte), unsafe; NewHex(string); and Parse(string) for
// creating UUIDs from existing data.
//
// The original version was from Krzysztof Kowalik <chris@nu7hat.ch>
// Unfortunately, that version was non compliant with RFC4122.
// I have since heavily redesigned it.
//
// The example code in the specification was also used as reference
// for design.
//
// Copyright (C) 2016 twinj@github.com  2014 MIT licence
package uuid

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/twinj/uuid/version"
	"hash"
	"regexp"
)

const (
	VariantNCS       uint8 = 0x00
	VariantRFC4122   uint8 = 0x80 // or and A0 if masked with 1F
	VariantMicrosoft uint8 = 0xC0
	VariantFuture    uint8 = 0xE0
)

type Domain uint8

const (
	DomainUser Domain = iota + 1
	DomainGroup
)

const (
	// Pattern used to parse string representation of the UUID.
	// Current one allows to parse string where only one opening
	// or closing bracket or any of the hyphens are optional.
	// It is only used to extract the main bytes to create a UUID,
	// so these imperfections are of no consequence.
	hexPattern = `^(urn\:uuid\:)?[\{\(\[]?([[:xdigit:]]{8})-?([[:xdigit:]]{4})-?([1-5][[:xdigit:]]{3})-?([[:xdigit:]]{4})-?([[:xdigit:]]{12})[\]\}\)]?$`
)

var (
	parseUUIDRegex = regexp.MustCompile(hexPattern)
)

func NewGenerator(
	fRandom func([]byte) (int, error),
	fNext func() Timestamp,
	fId func() Node) (generator *Generator) {
	generator = new(Generator)
	generator.Random = fRandom
	generator.Next = fNext
	generator.Id = fId
	return
}

// ******************************************************  UUID

// UUID is the common interface implemented by all UUIDs
type UUID interface {

	// Retrieves the UUID bytes from the underlying type data
	Bytes() []byte

	// Size is used where different implementations require different sizes.
	// Should return the number of bytes in the implementation.
	// Enables unmarshal and Bytes to screen for size
	Size() int

	// A UUID can be used as a Name within a namespace
	// Is simply just a String() string, method
	// Returns a formatted version of the UUID.
	UniqueName

	// Variant returns the UUID Variant
	// This will be one of the constants:
	// ReservedRFC4122,
	// ReservedMicrosoft,
	// ReservedFuture,
	// ReservedNCS.
	// This may behave differently across non RFC4122 UUIDs
	Variant() uint8

	// Version returns a version number of the algorithm used to generate the
	// UUID. This may may behave independently across non RFC4122 UUIDs
	Version() version.Version
}

// New creates a UUID from a slice of bytes.
func New(pData []byte) Uuid {
	o := array{}
	o.unmarshal(pData)
	return o[:]
}

// Creates a UUID from a hex string
// Will panic if hex string is invalid - will panic even with hyphens and brackets
// Expects a clean string use Parse otherwise.
func NewHex(pUuid string) Uuid {
	bytes, err := hex.DecodeString(pUuid)
	if err != nil {
		panic(err)
	}
	o := Uuid(bytes)
	return o
}

// Creates a UUID from a valid string representation.
// Accepts UUID string in following formats:
//		6ba7b8149dad11d180b400c04fd430c8
//		6ba7b814-9dad-11d1-80b4-00c04fd430c8
//		{6ba7b814-9dad-11d1-80b4-00c04fd430c8}
//		urn:uuid:6ba7b814-9dad-11d1-80b4-00c04fd430c8
//		[6ba7b814-9dad-11d1-80b4-00c04fd430c8]
//
func Parse(pUUID string) (UUID, error) {
	md := parseUUIDRegex.FindStringSubmatch(pUUID)
	if md == nil {
		return nil, errors.New("uuid.Parse: invalid string format this is probablt not a UUID")
	}
	return NewHex(md[2] + md[3] + md[4] + md[5] + md[6]), nil
}

// The RFC4122 implementation hashed the Namespace UUID in local byte order
// where they used a struct directly in the hash so its byte order compared to
// a slice was different. To get the same result as the specification we need
// to convert to this 'struct' representation
//
// type uuid struct {
// 	timeLow               uint32
// 	timeMid, timeHiAndVer uint16
// 	seqHiAndVar, seqLow   uint8
// 	id                    [6]uint8
// }
//
// var NameSpace_X500 = uuid{
// 	0x6ba7b814,
// 	0x9dad,
// 	0x11d1,
// 	0x80, 0xb4, []uint8{0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8},
// }
//
// var howTheStructLooksInMemory = [16]uint8{
// 	0x14, 0xb8, 0xa7, 0x6b, 0xad, 0x9d, 0xd1, 0x11, 0x80, 0xb4, 0x00, 0xc0,
// 	0x4f, 0xd4, 0x30, 0xc8,
// }
//
// var canonicalX500 = "6ba7b814-9dad-11d1-80b4-00c04fd430c8"
//
// var howItStoresInthisPackage = [16]uint8{
// 	0x6b, 0xa7, 0xb8, 0x14,
// 	0x9d, 0xad,
// 	0x11, 0xd1,
// 	0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8,
// }
//
// Since we work directly with bytes we need to convert those bytes to match
// how it was originally implemented to get the same computed hash values. Any
// string representation or natural representations in this package are
// expected to be in canonical order which is shown as rtl big-endian. It is
// possible to store all UUIDs in little endian order to avoid this issue,
// however, it is more efficient to have the bytes in the order they are
// represented to the world in so as to avoid overhead in marshalling and
// unmarshalling, printing and scanning. All overhead is in generation of an id
// for this reason. Some would argue that it would be better to store in little
// endian order so generating these would be faster from a benchmark
// perspective. I feel that this is only in service to computation in contrast
// to how the Ids are used and read.
func changeOrder(pName []byte) {
	groups := [][]byte{
		pName[:4], pName[4:6], pName[6:8],
	}
	for _, v := range groups {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
	}
}

func digest(pHash hash.Hash, pName []byte, pNames ...UniqueName) []byte {
	for _, v := range pNames {
		pName = append(pName, v.String()...)
	}
	pHash.Write(pName)
	return pHash.Sum(nil)
}

// **********************************************  UUID Names

// A UUID Name is a string which implements UniqueName
// which satisfies the Stringer interface. It is used to enable V3 and V5 UUIDs
// to use
type Name string

// Returns the name as a string. Satisfies the Stringer interface.
func (o Name) String() string {
	return string(o)
}

// UniqueName is a Stinger interface
// Made for easy passing of IPs, URLs, the several Address types,
// Buffers and any other type which implements Stringer
// string, []byte types and Hash sums will need to be cast to
// the Name type or some other type which implements
// UniqueName
type UniqueName interface {
	// Many go types implement this method for use with printing
	// Will convert the current type to its native string format
	String() string
}

// Compares whether each UUID is the same
func Equal(p1, p2 UUID) bool {
	return bytes.Equal(p1.Bytes(), p2.Bytes())
}

// Compare returns an integer comparing two UUIDs lexicographically.
// The result will be 0 if pId==pId2, -1 if pId < pId2, and +1 if pId > pId2.
// A nil argument is equivalent to the Nil UUID.
func Compare(pId, pId2 UUID) int {
	var b1, b2 []byte

	if pId == nil {
		b1 = []byte(Nil)
	} else {
		b1 = pId.Bytes()
	}

	if pId2 == nil {
		b2 = []byte(Nil)
	} else {
		b2 = pId2.Bytes()
	}

	return bytes.Compare(b1, b2)
}

// ***************************************************  Helpers

func variant(pVariant uint8) uint8 {
	switch pVariant & variantGet {
	case VariantRFC4122, 0xA0:
		return VariantRFC4122
	case VariantMicrosoft:
		return VariantMicrosoft
	case VariantFuture:
		return VariantFuture
	}
	return VariantNCS
}

func setVariant(pByte *byte, pVariant uint8) {
	switch pVariant {
	case VariantRFC4122:
		*pByte &= variantSet
	case VariantFuture, VariantMicrosoft:
		*pByte &= 0x1F
	case VariantNCS:
		*pByte &= 0x7F
	default:
		panic(errors.New("uuid.setVariant: invalid variant mask"))
	}
	*pByte |= pVariant
}
