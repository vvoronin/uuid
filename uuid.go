// This package provides RFC4122 and DCE 1.1 UUIDs.
//
// Use NewV1, NewV2, NewV3, NewV4, NewV5, for generating new UUIDs.
//
// Use New([]byte), NewHex(string), and Parse(string) for
// creating UUIDs from existing data.
//
// If you have a []byte you can simply cast it to the Uuid type.
//
// The original version was from Krzysztof Kowalik <chris@nu7hat.ch>
// Unfortunately, that version was non compliant with RFC4122.
//
// The package has since been redesigned.
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
	"regexp"
)

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
	Version() Version
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

// Compares whether each UUID is the same
func Equal(p1, p2 UUID) bool {
	return bytes.Equal(p1.Bytes(), p2.Bytes())
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
	return Uuid(bytes)
}

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

// Name is a string which implements UniqueName and satisfies the Stringer
// interface. V3 and V5 UUIDs use this for hashing values together to produce
// UUIDs based on a NameSpace.
type Name string

// String returns the uuid.Name as a string.
func (o Name) String() string {
	return string(o)
}

// UniqueName is a Stinger interface made for easy passing of any Stringer type
// into a Hashable UUID.
type UniqueName interface {
	// Many go types implement this method for use with printing
	// Will convert the current type to its native string format
	String() string
}

