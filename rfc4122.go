package uuid

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"encoding/hex"
)

// **************************************************** Namespaces

const (
	// The following standard UUIDs are for use with V3 or V5 UUIDs.
	// Note the big endian order for
	NamespaceDNS Name = "\x10\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NamespaceURL Name = "\x11\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NamespaceOID Name = "\x12\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
	NamespaceX500 Name = "\x14\xb8\xa7k\xad\x9d\xd1\x11\x80\xb4\x00\xc0O\xd40\xc8"
)

func mainTest() {
	fmt.Println("6ba7b8109dad11d180b400c04fd430c8")
	fmt.Println(len("6ba7b8109dad11d180b400c04fd430c8"))

	o, _ := hex.DecodeString("6ba7b8109dad11d180b400c04fd430c8")
	fmt.Printf("%#v\n", o)
	fmt.Printf("%#v\n", string(o))
	fmt.Printf("%#v\n", []byte(string(o)))

	o, _ = hex.DecodeString("6ba7b8119dad11d180b400c04fd430c8")
	fmt.Printf("%#v\n", o)
	fmt.Printf("%#v\n", string(o))

	o, _ = hex.DecodeString("6ba7b8119dad11d180b400c04fd430c8")
	fmt.Printf("%#v\n", o)
	fmt.Printf("%#v\n", string(o))

	o, _ = hex.DecodeString("6ba7b8129dad11d180b400c04fd430c8")
	fmt.Printf("%#v\n", o)
	fmt.Printf("%#v\n", string(o))

	o, _ = hex.DecodeString("6ba7b8149dad11d180b400c04fd430c8")
	fmt.Printf("%#v\n", string(o))
	fmt.Printf("%#v\n", []byte("\x6b\xa7\xb8\x10\x9d\xad\xd1\x11\x80\xb4\x00\xc0\x4f\xd4\x30\xc8"))

	h := md5.New()
	h.Write([]byte("\x10\xb8\xa7\x6b"))
	h.Write([]byte("\xad\x9d"))
	h.Write([]byte("\xd1\x11"))
	h.Write([]byte("\x80\xb4"))
	h.Write([]byte("\x00\xc0\x4f\xd4\x30\xc8"))
	h.Write([]byte("www.widgets.com"))
	fmt.Printf("%x\n", h.Sum(nil))

	h = md5.New()
	h.Write([]byte("\x6b\xa7\xb8\x10"))
	h.Write([]byte("\xad\x9d"))
	h.Write([]byte("\xd1\x11"))
	h.Write([]byte("\x80\xb4"))
	h.Write([]byte("\x00\xc0\x4f\xd4\x30\xc8"))
	h.Write([]byte("www.widgets.com"))
	fmt.Printf("%x\n", h.Sum(nil))

	h = md5.New()
	h.Write([]byte("\x6b\xa7\xb8\x10\x9d\xad\x11\xd1\x80\xb4\x00\xc0\x4f\xd4\x30\xc8"))
	h.Write([]byte("www.widgets.com"))
	fmt.Printf("%x\n", h.Sum(nil))

	// e902893a-9d22-3c7e-a7b8-d6e313b71d9f
	// 3a8902e9-229d-7e7c-e7b8-d6e313b71d9f

	h = md5.New()
	h.Write([]byte("k\xa7\xb8\x10\x9d\xad\x11р\xb4\x00\xc0O\xd40\xc8"))
	h.Write([]byte("www.widgets.com"))
	fmt.Printf("%x\n", h.Sum(nil))
}

var (
	namespaceDNS UUID = FromName(NamespaceDNS)
	namespaceURL UUID = FromName(NamespaceURL)
	namespaceOID UUID = FromName(NamespaceOID)
	namespaceX500 UUID = FromName(NamespaceX500)
)

var (
	generator *Generator
)

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
// based on a 60 bit timestamp, node id and POSIX UID
func NewV2(pDomain Domain) UUID {
	return generator.NewV2(pDomain)
}

// Generates a new RFC4122 version 3 UUID
// Based on the MD5 hash of a namespace UUID and
// any type which implements the UniqueName interface for the name.
// For strings and slices cast to a Name type
func NewV3(pNamespace Name, pNames ...UniqueName) UUID {
	n := digest(md5.New(), pNamespace, pNames...)
	o := fromName(n)
	o.setRFC4122Version(3)
	return &o
}

// Generates a new RFC4122 version 4 UUID
// A cryptographically secure random UUID.
func NewV4() UUID {
	o := make(array, length)
	rand.Read(o)
	o.setRFC4122Version(4)
	return &o
}

// NewV5 generates an RFC4122 version 5 UUID
// based on the SHA-1 hash of a namespace
// UUID and a unique name.
func NewV5(pNamespace Name, pNames ...UniqueName) UUID {
	n := digest(sha1.New(), pNamespace, pNames...)
	o := fromName(n)
	o.setRFC4122Version(5)
	return &o
}
