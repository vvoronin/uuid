package uuid

/****************
 * Date: 16/02/14
 * Time: 11:29 AM
 ***************/

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

const (
	generate = 10000
)

var (
	namespaces = make(map[UUID]string)
)

func init() {
	namespaces[NameSpaceX500] = "6ba7b814-9dad-11d1-80b4-00c04fd430c8"
	namespaces[NameSpaceOID] = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	namespaces[NameSpaceURL] = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"
	namespaces[NameSpaceDNS] = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
}

func TestNewV1(t *testing.T) {
	generator.Do(generator.init)
	u := NewV1()
	assert.Equal(t, One, u.Version(), "Expected correct version")
	assert.Equal(t, VariantRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")
}

func TestNewV2(t *testing.T) {
	u := NewV2(DomainGroup)

	assert.Equal(t, Two, u.Version(), "Expected correct version")
	assert.Equal(t, VariantRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")
}

func TestNewV3(t *testing.T) {
	u := NewV3(NameSpaceURL, goLang)

	assert.Equal(t, Three, u.Version(), "Expected correct version")
	assert.Equal(t, VariantRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")

	ur, _ := url.Parse(string(goLang))

	// Same NS same name MUST be equal
	u2 := NewV3(NameSpaceURL, ur)
	assert.Equal(t, u, u2, "Expected UUIDs generated with same namespace and name to equal")

	// Different NS same name MUST NOT be equal
	u3 := NewV3(NameSpaceDNS, ur)
	assert.NotEqual(t, u, u3, "Expected UUIDs generated with different namespace and same name to be different")

	// Same NS different name MUST NOT be equal
	u4 := NewV3(NameSpaceURL, u)
	assert.NotEqual(t, u, u4, "Expected UUIDs generated with the same namespace and different names to be different")

	ids := []UUID{
		u, u2, u3, u4,
	}

	for j, id := range ids {
		i := NewV3(NameSpaceURL, Name(string(j)), id)
		assert.NotEqual(t, id, i, "Expected UUIDs generated with the same namespace and different names to be different")
	}

	u = NewV3(NameSpaceDNS, Name("www.example.com"))
	assert.Equal(t, "5df41881-3aed-3515-88a7-2f4a814cf09e", u.String())

	u = NewV3(NameSpaceDNS, Name("python.org"))
	assert.Equal(t, "6fa459ea-ee8a-3ca4-894e-db77e160355e", u.String())
}

func TestNewV4(t *testing.T) {
	u := NewV4()
	assert.Equal(t, Four, u.Version(), "Expected correct version")
	assert.Equal(t, VariantRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")
}

func TestNewV5(t *testing.T) {
	u := NewV5(NameSpaceURL, goLang)

	assert.Equal(t, Five, u.Version(), "Expected correct version")
	assert.Equal(t, VariantRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")

	ur, _ := url.Parse(string(goLang))

	// Same NS same name MUST be equal
	u2 := NewV5(NameSpaceURL, ur)
	assert.Equal(t, u, u2, "Expected UUIDs generated with same namespace and name to equal")

	// Different NS same name MUST NOT be equal
	u3 := NewV5(NameSpaceDNS, ur)
	assert.NotEqual(t, u, u3, "Expected UUIDs generated with different namespace and same name to be different")

	// Same NS different name MUST NOT be equal
	u4 := NewV5(NameSpaceURL, u)
	assert.NotEqual(t, u, u4, "Expected UUIDs generated with the same namespace and different names to be different")

	ids := []UUID{
		u, u2, u3, u4,
	}

	for j, id := range ids {
		i := NewV5(NameSpaceURL, Name(string(j)), id)
		assert.NotEqual(t, i, id, "Expected UUIDs generated with the same namespace and different names to be different")
	}

	u = NewV5(NameSpaceDNS, Name("python.org"))
	assert.Equal(t, "886313e1-3b8a-5372-9b90-0c9aee199e5d", u.String())
}

var printIt = false

func printer(pId Uuid) {
	if printIt {
		fmt.Println(pId)
	}
}

func TestUUID_NewV1Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		printer(NewV1())
	}
}

func TestUUID_NewV3Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		printer(NewV3(NameSpaceDNS, goLang, Name(string(i))))
	}
}

func TestUUID_NewV4Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		printer(NewV4())
	}
}

func TestUUID_NewV5Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		printer(NewV5(NameSpaceDNS, goLang, Name(string(i))))
	}
}

func Test_EachIsUnique(t *testing.T) {

	// Run half way through to avoid running within default resolution only

	spin := int(defaultSpinResolution / 2)

	for i := 0; i < spin; i++ {
		NewV1()
	}

	s := int(defaultSpinResolution)

	ids := make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV1()
		ids[i] = u
		for j := 0; j < i; j++ {
			if b := assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V1 UUID"); !b {
				break
			}
		}
	}
	//ids = make([]UUID, s)
	//for i := 0; i < s; i++ {
	//	u := NewV2(DomainGroup)
	//	ids[i] = u
	//	for j := 0; j < i; j++ {
	//		assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V2 UUID")
	//	}
	//}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV3(NameSpaceDNS, Name(string(i)), goLang)
		ids[i] = u
		for j := 0; j < i; j++ {
			if b := assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V3 UUID"); !b {
				break
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV4()
		ids[i] = u
		for j := 0; j < i; j++ {
			if b := assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V4 UUID"); !b {
				break
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV5(NameSpaceDNS, Name(string(i)), goLang)
		ids[i] = u
		for j := 0; j < i; j++ {
			if b := assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V5 UUID"); !b {
				break
			}
		}
	}
}

func Test_NameSpaceUUIDs(t *testing.T) {
	for k, v := range namespaces {

		arrayId, _ := Parse(v)
		uuidId := array{}
		uuidId.unmarshal(arrayId.Bytes())
		assert.Equal(t, v, arrayId.String())
		assert.Equal(t, v, k.String())
	}
}

func TestNewV12(t *testing.T) {
	id := array{}

	makeUuid(&id,
		0x6ba7b810,
		0x9dad,
		0x11d1,
		0x80b4,
		[]byte{0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8})

	assert.Equal(t, id[:], NameSpaceDNS.Bytes())
	fmt.Println(Uuid(id[:]))
}
