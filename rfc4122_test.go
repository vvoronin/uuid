package uuid

/****************
 * Date: 16/02/14
 * Time: 11:29 AM
 ***************/

import (
	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid/version"
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

func TestV1(t *testing.T) {
	u := NewV1()

	assert.Equal(t, version.One, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")
}

func TestV2(t *testing.T) {
	u := NewV2(DomainGroup)

	assert.Equal(t, version.Two, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")
}

func TestV3(t *testing.T) {
	u := NewV3(NameSpaceURL, goLang)

	assert.Equal(t, version.Three, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
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

	u = NewV3(NameSpaceDNS, Name("www.widgets.com"))

	assert.Equal(t, "e902893a-9d22-3c7e-a7b8-d6e313b71d9f", u.String())
}

func TestV4(t *testing.T) {
	u := NewV4()
	assert.Equal(t, version.Four, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")
}

func TestV5(t *testing.T) {
	u := NewV5(NameSpaceURL, goLang)

	assert.Equal(t, version.Five, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
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

}

func TestUUID_NewV1Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV1()
	}
}

func TestUUID_NewV3Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV3(NameSpaceDNS, goLang)
	}
}

func TestUUID_NewV4Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV4()
	}
}

func TestUUID_NewV5Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV5(NameSpaceDNS, goLang)
	}
}

func Test_EachIsUnique(t *testing.T) {
	s := 20
	ids := make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV1()
		ids[i] = u
		for j := 0; j < i; j++ {
			assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V1 UUID")
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
			assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V3 UUID")

		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV4()
		ids[i] = u
		for j := 0; j < i; j++ {
			assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V4 UUID")
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV5(NameSpaceDNS, Name(string(i)), goLang)
		ids[i] = u
		for j := 0; j < i; j++ {
			assert.NotEqual(t, u.String(), ids[j].String(), "Should not create the same V5 UUID")
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
