package uuid_test

import (
	"fmt"
	"github.com/twinj/uuid"
	"github.com/twinj/uuid/savers"
	"time"
)

func Example() {
	saver := new(savers.FileSystemSaver)
	saver.Report = true
	saver.Duration = time.Second * 3

	// Run before any v1 or v2 UUIDs to ensure the savers takes
	uuid.RegisterSaver(saver)

	u1 := uuid.NewV1()
	fmt.Printf("version %d variant %x: %s\n", u1.Version(), u1.Variant(), u1)

	uP, _ := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	uP = uuid.PromoteToNameSpace(uP)
	u3 := uuid.NewV3(uP.(uuid.NameSpace), uuid.Name("test"))

	u4 := uuid.NewV4()
	fmt.Printf("version %d variant %x: %s\n", u4.Version(), u4.Variant(), u4)

	u5 := uuid.NewV5(uuid.NameSpaceURL, uuid.Name("test"))

	if uuid.Equal(u1, u3) {
		fmt.Println("Will never happen")
	}

	fmt.Println(uuid.Sprintf(uuid.CurlyHyphen, u5))

	uuid.SwitchFormat(uuid.BracketHyphen)
}

func ExampleNewV1() {
	u1 := uuid.NewV1()
	fmt.Printf("version %d variant %s: %s\n", u1.Version(), u1.Variant(), u1)
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

func ExampleSprintf() {
	u4 := uuid.NewV4()
	fmt.Println(uuid.Sprintf(uuid.CurlyHyphen, u4))
}

func ExampleSwitchFormat() {
	uuid.SwitchFormat(uuid.BracketHyphen)
	u4 := uuid.NewV4()
	fmt.Printf("version %d variant %x: %s\n", u4.Version(), u4.Variant(), u4)
}
