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
	// nodeID is the default Namespace node
	nodeId = []byte{
		// 00.192.79.212.48.200
		0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8,
	}
	// The following standard UUIDs are for use with V3 or V5 UUIDs.
	NamespaceDNS  = &uuid{0x6ba7b810, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}
	NamespaceURL  = &uuid{0x6ba7b811, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}
	NamespaceOID  = &uuid{0x6ba7b812, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}
	NamespaceX500 = &uuid{0x6ba7b814, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}

	state = generator{}
)


// NewV1 will generate a new RFC4122 version 1 UUID
func NewV1() UUID {
	store := state.read()

	o := new(uuid)

	o.timeLow = uint32(store.Timestamp & 0xffffffff)
	o.timeMid = uint16((store.Timestamp >> 32) & 0xffff)
	o.timeHiAndVersion = uint16((store.Timestamp >> 48) & 0x0fff)
	o.timeHiAndVersion |= uint16(1 << 12)
	o.sequenceLow = byte(store.Sequence & 0xff)
	o.sequenceHiAndVariant = byte((store.Sequence & 0x3f00) >> 8)
	o.sequenceHiAndVariant |= ReservedRFC4122
	o.node = make([]byte, len(store.Node))
	copy(o.node[:], store.Node)
	o.size = length

	return o
}

func NewV2(pDomain DCEDomain) UUID {

	//now, sequence, node := state.read()

	o := new(uuid)

	//switch pDomain {
	//	case DomainPerson:
	//		binary.BigEndian.PutUint32(u[0:], posixUID)
	//	case DomainGroup:
	//	binary.BigEndian.PutUint32(u[0:], posixGID)
	//}
	//
	//o.timeLow = uint32(now & 0xffffffff)
	//o.timeMid = uint16((now >> 32) & 0xffff)
	//o.timeHiAndVersion = uint16((now >> 48) & 0x0fff)
	//o.timeHiAndVersion |= uint16(1 << 12)
	//o.sequenceLow = byte(sequence & 0xff)
	//o.sequenceHiAndVariant = byte((sequence & 0x3f00) >> 8)
	//o.sequenceHiAndVariant |= ReservedRFC4122
	//o.node = make([]byte, len(node))
	//copy(o.node[:], node)
	//o.size = length

	return o
}

// NewV3 will generate a new RFC4122 version 3 UUID
// V3 is based on the MD5 hash of a namespace identifier UUID and
// any type which implements the UniqueName interface for the name.
// For strings and slices cast to a Name type
func NewV3(pNs UUID, pName UniqueName) UUID {
	o := new(array)
	// Set all bits to MD5 hash generated from namespace and name.
	Digest(o, pNs, pName, md5.New())
	o.setRFC4122Variant()
	o.setVersion(3)
	return o
}

// NewV4 will generate a new RFC4122 version 4 UUID
// A cryptographically secure random UUID.
func NewV4() UUID {
	o := new(array)
	// Read random values (or pseudo-randomly) into Array type.
	_, err := rand.Read(o[:length])
	if err != nil {
		panic(err)
	}
	o.setRFC4122Variant()
	o.setVersion(4)

	return o
}

// NewV5 will generate a new RFC4122 version 5 UUID
// Generate a UUID based on the SHA-1 hash of a namespace
// identifier and a name.
func NewV5(pNs UUID, pName UniqueName) UUID {
	o := new(array)
	Digest(o, pNs, pName, sha1.New())
	o.setRFC4122Variant()
	o.setVersion(5)
	return o
}
