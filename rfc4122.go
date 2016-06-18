package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"hash"
	"log"
)

const (
	Nil uuid = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"

	// The following standard UUIDs are for use with V3 or V5 UUIDs.
	// Note the big endian order for each octet set - this is to ensure compliant hash outputs
	NameSpaceDNS  uuid = "k\xa7\xb8\x10\x9d\xad\x11р\xb4\x00\xc0O\xd40\xc8"
	NameSpaceURL  uuid = "k\xa7\xb8\x11\x9d\xad\x11р\xb4\x00\xc0O\xd40\xc8"
	NameSpaceOID  uuid = "k\xa7\xb8\x12\x9d\xad\x11р\xb4\x00\xc0O\xd40\xc8"
	NameSpaceX500 uuid = "k\xa7\xb8\x14\x9d\xad\x11р\xb4\x00\xc0O\xd40\xc8"
)

type Domain uint8

const (
	DomainUser Domain = iota + 1
	DomainGroup
)

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
func NewV3(pNamespace UUID, pNames ...UniqueName) Uuid {
	o := array{}
	copy(o[:], digest(md5.New(), pNamespace.Bytes(), pNames...))
	o.setRFC4122Version(3)
	return o[:]
}

// NewV4 generates a new RFC4122 version 4 UUID a cryptographically secure
// random UUID.
func NewV4() Uuid {
	o, err := v4()
	if err == nil {
		return o[:]
	}
	generator.err = err
	log.Printf("uuid.V4: There was an error getting random bytes [%s]\n", err)
	if ok := generator.HandleError(err); ok {
		o, err = v4()
		if err == nil {
			return o[:]
		}
		generator.err = err
	}
	return nil
}

func v4() (o array, err error) {
	generator.err = nil
	_, err = generator.Random(o[:])
	o.setRFC4122Version(4)
	return
}

// NewV5 generates an RFC4122 version 5 UUID based on the SHA-1 hash of a
// namespace UUID and a unique name.
func NewV5(pNamespace UUID, pNames ...UniqueName) Uuid {
	o := array{}
	copy(o[:], digest(sha1.New(), pNamespace.Bytes(), pNames...))
	o.setRFC4122Version(5)
	return o[:]
}

func digest(pHash hash.Hash, pName []byte, pNames ...UniqueName) []byte {
	for _, v := range pNames {
		pName = append(pName, v.String()...)
	}
	pHash.Write(pName)
	return pHash.Sum(nil)
}
