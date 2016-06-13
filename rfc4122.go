package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"log"
	"hash"
)

const (
	Nil uuid = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"

	// The following standard UUIDs are for use with V3 or V5 UUIDs.
	// Note the big endian order for each octet set - this is to ensure compliant hash outputs
	NameSpaceDNS  NameSpace = "\x10\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NameSpaceURL  NameSpace = "\x11\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NameSpaceOID  NameSpace = "\x12\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NameSpaceX500 NameSpace = "\x14\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
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
// var howItStoresInThisPackage = [16]uint8{
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


