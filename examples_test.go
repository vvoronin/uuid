package uuid_test

import (
	"fmt"
	"github.com/twinj/uuid"
	"github.com/twinj/uuid/savers"
	"time"
	"net/url"
)

func Example() {
	saver := new(savers.FileSystemSaver)
	saver.Report = true
	saver.Duration = time.Second * 3

	// Run before any v1 or v2 UUIDs to ensure the savers takes
	uuid.RegisterSaver(saver)

	up, _ := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	fmt.Printf("version %d variant %x: %s\n", up.Version(), up.Variant(), up)

	uuid.New(up.Bytes())

	u1 := uuid.NewV1()
	fmt.Printf("version %d variant %x: %s\n", u1.Version(), u1.Variant(), u1)

	u4 := uuid.NewV4()
	fmt.Printf("version %d variant %x: %s\n", u4.Version(), u4.Variant(), u4)

	newNameSpace := uuid.PromoteToNameSpace(u1)
	u3 := uuid.NewV3(newNameSpace, u4)

	url, _ := url.Parse("www.widgets.com")

	u5 := uuid.NewV5(uuid.NameSpaceURL, url)

	if uuid.Equal(u1, u3) {
		fmt.Println("Will never happen")
	}

	if uuid.Compare(uuid.NameSpaceDNS, uuid.NameSpaceDNS) == 0 {
		fmt.Println("They are equal")
	}

	// Default Format is Canonical
	fmt.Println(uuid.Formatter(u5, uuid.CanonicalCurly))

	uuid.SwitchFormat(uuid.CanonicalBracket)
}

func ExampleNewV1() {
	u1 := uuid.NewV1()
	fmt.Printf("version %d variant %s: %s\n", u1.Version(), u1.Variant(), u1)
}

func ExampleNewV2() {
	u2 := uuid.NewV2(uuid.DomainUser)
	fmt.Printf("version %d variant %s: %s\n", u2.Version(), u2.Variant(), u2)
}

func ExampleNewV3() {
	u, _ := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	n := uuid.PromoteToNameSpace(u)
	u3 := uuid.NewV3(n, uuid.Name("test"))
	fmt.Printf("version %d variant %x: %s\n", u3.Version(), u3.Variant(), u3)
}

func ExampleNewV4() {
	u4 := uuid.NewV4()
	fmt.Printf("version %d variant %x: %s\n", u4.Version(), u4.Variant(), u4)
}

func ExampleNewV5() {
	u5 := uuid.NewV5(uuid.NameSpaceURL, uuid.Name("test"))
	fmt.Printf("version %d variant %x: %s\n", u5.Version(), u5.Variant(), u5)
}

func ExampleParse() {
	u, err := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(u)
}

func ExampleRegisterSaver() {
	saver := new(savers.FileSystemSaver)
	saver.Report = true
	saver.Duration = 3 * time.Second

	// Run before any v1 or v2 UUIDs to ensure the savers takes
	uuid.RegisterSaver(saver)
	u1 := uuid.NewV1()
	fmt.Printf("version %d variant %x: %s\n", u1.Version(), u1.Variant(), u1)
}

func ExampleFormatter() {
	u4 := uuid.NewV4()
	fmt.Println(uuid.Formatter(u4, uuid.CanonicalCurly))
}

func ExampleSwitchFormat() {
	uuid.SwitchFormat(uuid.CanonicalBracket)
	u4 := uuid.NewV4()
	fmt.Printf("version %d variant %x: %s\n", u4.Version(), u4.Variant(), u4)
}
