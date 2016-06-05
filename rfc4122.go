package uuid

/***************
 * Date: 14/02/14
 * Time: 7:44 PM
 ***************/

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
)

// ****************************************************

const (
	length = 16

	// 3f used by RFC4122 although 1f works for all
	variantSet = 0x3f

	// rather than using 0xc0 we use 0xe0 to retrieve the variant
	// The result is the same for all other variants
	// 0x80 and 0xa0 are used to identify RFC4122 compliance
	variantGet = 0xe0
)

var (
	// The following standard UUIDs are for use with V3 or V5 UUIDs.
	NamespaceDNS  UUID = NewHex("6ba7b8109dad11d180b400c04fd430c8")
	NamespaceURL  UUID = NewHex("6ba7b8119dad11d180b400c04fd430c8")
	NamespaceOID  UUID = NewHex("6ba7b8129dad11d180b400c04fd430c8")
	NamespaceX500 UUID = NewHex("6ba7b8149dad11d180b400c04fd430c8")

	generator *Generator
)

// ****************************************************

func init() {
	registerDefaultGenerator()
}

func registerDefaultGenerator() {
	generator = newGenerator(
		(&spinner{
			Resolution: 512,
			Timestamp:  Now(),
			Count:      0,
		}).next,
		findFirstHardwareAddress,
		CleanHyphen)
}

// Generate a new RFC4122 version 1 UUID
// based on a 60 bit timestamp and node id
func NewV1() UUID {
	return generator.NewV1()
}

// Generate a new DCE Security version UUID
// based on a 60 bit timestamp, node id and POSIX UID or GUID
func NewV2(pDomain DCEDomain) UUID {
	return generator.NewV2(pDomain)
}

// Generates a new RFC4122 version 3 UUID
// Based on the MD5 hash of a namespace UUID and
// any type which implements the UniqueName interface for the name.
// For strings and slices cast to a Name type
func NewV3(pNs UUID, pName UniqueName) UUID {
	// Benchmarks faster
	o := &array{}

	h := md5.New()
	h.Write(pNs.Bytes())
	h.Write([]byte(pName.String()))
	copy(o[:], h.Sum(nil))

	o[versionIndex] &= 0x0f
	o[versionIndex] |= 3 << 4
	o[variantIndex] &= variantSet
	o[variantIndex] |= ReservedRFC4122
	return o
}

// Generates a new RFC4122 version 3 UUID
// Based on the MD5 hash of a namespace UUID and
// any type which implements the UniqueName interface for the name.
// For strings and slices cast to a Name type
func V3(pNs []byte, pName string) UUID {
	o := array{}
	// Set all bits to MD5 hash generated from namespace and name.
	h := md5.New()
	h.Write(pNs)
	h.Write([]byte(pName))
	copy(o[:], h.Sum(nil))

	o[versionIndex] &= 0x0f
	o[versionIndex] |= 3 << 4
	o[variantIndex] &= variantSet
	o[variantIndex] |= ReservedRFC4122
	return &o
}

// Generates a new RFC4122 version 4 UUID
// A cryptographically secure random UUID.
func NewV4() UUID {
	// Benchmarks faster
	o := &array{}

	// Read random values (or pseudo-random) into array type.
	rand.Read(o[:length])
	o.setVersion(4)
	o.setRFC4122Variant()
	return o
}

// Generates a new RFC4122 version 5 UUID
// based on the SHA-1 hash of a namespace
// UUID and a unique name.
func NewV5(pNs UUID, pName UniqueName) UUID {
	// Benchmarks faster
	o := &array{}
	digest(o, pNs, pName, sha1.New())
	o.setVersion(5)
	o.setRFC4122Variant()
	return o
}
