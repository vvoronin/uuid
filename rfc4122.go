package uuid

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"log"
)

// **************************************************** Namespaces

const (
	Nil uuid = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"

	// The following standard UUIDs are for use with V3 or V5 UUIDs.
	// Note the big endian order for each octet set - this is to ensure compliant hash outputs
	NameSpaceDNS  NameSpace = "\x10\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NameSpaceURL  NameSpace = "\x11\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NameSpaceOID  NameSpace = "\x12\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NameSpaceX500 NameSpace = "\x14\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
)

var (
	generator *Generator
)

func init() {
	registerDefaultGenerator()
}

func registerDefaultGenerator() {
	generator = NewGenerator(
		rand.Read,
		(&spinner{
			Resolution: 512,
			Timestamp:  Now(),
			Count:      0,
		}).next,
		findFirstHardwareAddress)
}

// NewV1 generates a new RFC4122 version 1 UUID based on a 60 bit timestamp and
// node ID.
func NewV1() Uuid {
	return generator.NewV1()
}

// NewV2 generates a new DCE Security version UUID based on a 60 bit timestamp,
// node id and POSIX UID.
func NewV2(pDomain Domain) Uuid {
	return generator.NewV2(pDomain)
}

// NewV3 generates a new RFC4122 version 3 UUID based on the MD5 hash on a
// namespace UUID and any type which implements the UniqueName interface
// for the name. For strings and slices cast to a Name type
func NewV3(pNamespace NameSpace, pNames ...UniqueName) Uuid {
	o := array{}
	copy(o[:], digest(md5.New(), []byte(pNamespace), pNames...))
	changeOrder(o[:])
	o.setRFC4122Version(3)
	return o[:]
}

// NewV4 generates a new RFC4122 version 4 UUID a cryptographically secure
// random UUID.
func NewV4() Uuid {
	o := array{}
	_, err := generator.Random(o[:])
	if err == nil {
		o.setRFC4122Version(4)
		return o[:]
	}
	generator.err = err
	log.Printf("uuid.V4: There was an error getting random bytes [%s]\n", err)
	return nil
}

// NewV5 generates an RFC4122 version 5 UUID based on the SHA-1 hash of a
// namespace UUID and a unique name.
func NewV5(pNamespace NameSpace, pNames ...UniqueName) Uuid {
	o := array{}
	copy(o[:], digest(sha1.New(), []byte(pNamespace), pNames...))
	changeOrder(o[:])
	o.setRFC4122Version(5)
	return o[:]
}

func PromoteToNameSpace(pId UUID) NameSpace {
	if v, ok := pId.(NameSpace); ok {
		return v
	}
	o := array{}
	o.unmarshal(pId.Bytes())
	changeOrder(o[:])
	return NameSpace(o[:])
}
