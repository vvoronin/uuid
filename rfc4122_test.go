package uuid

/****************
 * Date: 16/02/14
 * Time: 11:29 AM
 ***************/

import (
	"net/url"
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	goLang Name = "https://google.com/golang.org?q=golang"
)

const (
	generate = 10000
)

func TestNewV1(t *testing.T) {
	u := NewV1()

	assert.Equal(t, 1, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")
}

func TestNewV2(t *testing.T) {

}

func TestNewV3(t *testing.T) {
	u := NewV3(NamespaceURL, goLang)

	assert.Equal(t, 3, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")

	ur, _ := url.Parse(string(goLang))

	// Same NS same name MUST be equal
	u2 := NewV3(NamespaceURL, ur)
	assert.Equal(t, u, u2, "Expected UUIDs generated with same namespace and name to equal")

	// Different NS same name MUST NOT be equal
	u3 := NewV3(NamespaceDNS, ur)
	assert.NotEqual(t, u, u3, "Expected UUIDs generated with different namespace and same name to be different")

	// Same NS different name MUST NOT be equal
	u4 := NewV3(NamespaceURL, u)
	assert.NotEqual(t, u, u4, "Expected UUIDs generated with the same namespace and different names to be different")

	ids := []UUID{
		u, u2, u3, u4,
	}

	for j, id := range ids {
		i := NewV3(NamespaceURL, NewName(string(j), id))
		assert.NotEqual(t, id, i, "Expected UUIDs generated with the same namespace and different names to be different")
	}
}

func TestNewV4(t *testing.T) {
	u := NewV4()

	assert.Equal(t, 4, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")

}

func TestNewV5(t *testing.T) {
	u := NewV5(NamespaceURL, goLang)

	assert.Equal(t, 5, u.Version(), "Expected correct version")
	assert.Equal(t, ReservedRFC4122, u.Variant(), "Expected correct variant")
	assert.True(t, parseUUIDRegex.MatchString(u.String()), "Expected string representation to be valid")

	ur, _ := url.Parse(string(goLang))

	// Same NS same name MUST be equal
	u2 := NewV5(NamespaceURL, ur)
	assert.Equal(t, u, u2, "Expected UUIDs generated with same namespace and name to equal")


	// Different NS same name MUST NOT be equal
	u3 := NewV5(NamespaceDNS, ur)
	assert.NotEqual(t, u, u3, "Expected UUIDs generated with different namespace and same name to be different")

	// Same NS different name MUST NOT be equal
	u4 := NewV5(NamespaceURL, u)
	assert.NotEqual(t, u, u4, "Expected UUIDs generated with the same namespace and different names to be different")

	ids := []UUID{
		u, u2, u3, u4,
	}

	for j, id := range ids {
		i := NewV5(NamespaceURL, NewName(string(j), id))
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
		NewV3(NamespaceDNS, goLang)
	}
}

func TestUUID_NewV4Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV4()
	}
}

func TestUUID_NewV5Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV5(NamespaceDNS, goLang)
	}
}

// A small test to test uniqueness across all UUIDs created
func Test_EachIsUnique(t *testing.T) {
	s := 1000
	ids := make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV1()
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V1 UUID", u, ids[j])
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV3(NamespaceDNS, NewName(string(i), Name(goLang)))
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V3 UUID", u, ids[j])
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV4()
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V4 UUID", u, ids[j])
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV5(NamespaceDNS, NewName(string(i), Name(goLang)))
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V5 UUID", u, ids[j])
			}
		}
	}
}

func Test_Rfc4122(t *testing.T) {
	id, err := Parse("7d444840-9dc0-11d1-b245-5ffdce74fad2")

	assert.Nil(t, err, "There should be no error")
	assert.True(t, Equal(id, id), "The two ids should be equal")
	assert.False(t, Equal(id, NamespaceDNS), "The two ids should not be equal")

	v3 := NewV3(NamespaceDNS, Name("www.widgets.com"))

	assert.Equal(t, "e902893a-9d22-3c7e-a7b8-d6e313b71d9f", v3.String(), "The strings shoud be the same")
}

var namespaces = make(map[UUID]string)

func init() {
	namespaces[NamespaceX500] = "6ba7b814-9dad-11d1-80b4-00c04fd430c8"
	namespaces[NamespaceOID] = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	namespaces[NamespaceURL] = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"
	namespaces[NamespaceDNS] = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
}

func TestNameSpaceUUIDs(t *testing.T) {
	for k, v := range namespaces {

		arrayId, _ := Parse(v)

		uuidId := new(uuid)
		uuidId.UnmarshalBinary(arrayId.Bytes())
		assert.Equal(t, v, arrayId.String())
		assert.Equal(t, v, k.String())
	}
}