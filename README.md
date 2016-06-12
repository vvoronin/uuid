Go UUID implementation
========================

[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/twinj/uuid/master/LICENSE)
[![GoDoc](http://godoc.org/github.com/twinj/uuid?status.png)](http://godoc.org/github.com/twinj/uuid)
[![Build Status](https://ci.appveyor.com/api/projects/status/github/twinj/uuid?branch=master&svg=true)](https://ci.appveyor.com/project/twinj/uuid)
[![Build Status](https://travis-ci.org/twinj/uuid.png?branch=master)](https://travis-ci.org/twinj/uuid)
[![Coverage Status](https://coveralls.io/repos/github/twinj/uuid/badge.svg?branch=master)](https://coveralls.io/github/twinj/uuid?branch=master)

**This project is currently pre 1.0.0**


This package provides RFC 4122 and DCE 1.1 compliant UUIDs.
It will generate the following:

* Version 1: based on a Timestamp and MAC address as Node id
* Version 2: coming
* Version 3: based on MD5 hash
* Version 4: based on cryptographically secure random numbers
* Version 5: based on SHA-1 hash
* Your own implementations

Functions NewV1, NewV2, NewV3, NewV4, NewV5, New, NewHex and Parse() for generating version 1, 2, 3, 4
and 5 UUIDs

# Requirements

Any supported version of Go.

# Design considerations

* Ensure UUIDs are unique across a use case
    Proper test coverage has determined thant the UUID timestamp spinner works correctly, cross multiple clock resolutions
    The generator produces timestamps that roll out sequentially and will only modify the sequence on rare circumstances
    It is highly recommended that you register a Saver if you use V1 or V2 UUIDs as it will ensure a higher probability
    of uniqueness.

    Example V1 output:
    5fb1a280-30f0-11e6-9614-005056c00001
    5fb1a281-30f0-11e6-9614-005056c00001
    5fb1a282-30f0-11e6-9614-005056c00001
    5fb1a283-30f0-11e6-9614-005056c00001
    5fb1a284-30f0-11e6-9614-005056c00001
    5fb1a285-30f0-11e6-9614-005056c00001
    5fb1a286-30f0-11e6-9614-005056c00001
    5fb1a287-30f0-11e6-9614-005056c00001
    5fb1a288-30f0-11e6-9614-005056c00001
    5fb1a289-30f0-11e6-9614-005056c00001
    5fb1a28a-30f0-11e6-9614-005056c00001
    5fb1a28b-30f0-11e6-9614-005056c00001
    5fb1a28c-30f0-11e6-9614-005056c00001
    5fb1a28d-30f0-11e6-9614-005056c00001
    5fb1a28e-30f0-11e6-9614-005056c00001
    5fb1a28f-30f0-11e6-9614-005056c00001
    5fb1a290-30f0-11e6-9614-005056c00001

* Generator should work on all app servers.
    No Os locking threads or file system dependant storage
    Saver interface exists for the user to provide their own Saver implementations
    for V1 and V2 UUIDs. The interface could theoretically be applied to your own UUID implementation.
    Have provided a savers which works on a standard OS environment.
* Allow user implementations

# Future considerations

* length and format of UUID should not be an issue
* using new cryptographic technology should not be an issue
* improve support for sequential UUIDs merged with cryptographic nodes

# Recent Changes

* Improved builds and 100% test coverage
* Library overhaul to cleanup exports that are not useful for a user
* Improved file system Saver interface, breaking changes.
    To use a savers make sure you pass it in via the uuid.SetupSaver(Saver) method before a UUID is generated, so as to take affect.
* Removed use of OS Thread locking and runtime package requirement
* Changed String() output to CleanHyphen to match the canonical standard
* Removed default non volatile store and replaced with Saver interface
* Added formatting support for user defined formats
* Variant type bits are now set correctly
* Variant type can now be retrieved more efficiently
* New tests for variant setting to confirm correctness
* New tests added to confirm proper version setting

## Installation

Use the `go` tool:

	$ go get github.com/twinj/uuid

## Usage

See [documentation and examples](http://godoc.org/github.com/twinj/uuid)
for more information.

	saver := new(savers.FileSystemSaver)
	saver.Report = true
	saver.Duration = time.Second * 3

	// Run before any v1 or v2 UUIDs to ensure the savers takes
	uuid.RegisterSaver(saver)

	uP, _ := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	u1 := uuid.NewV1()
	fmt.Printf("version %d variant %x: %s\n", u1.Version(), u1.Variant(), u1)

	u4 := uuid.NewV4()
	fmt.Printf("version %d variant %x: %s\n", u4.Version(), u4.Variant(), u4)

	newNameSpace := uuid.PromoteToNameSpace(u1)
	u3 := uuid.NewV3(newNameSpace, u4)

	u5 := uuid.NewV5(uuid.NameSpaceURL, url.Parse("www.widgets.com"))

	if uuid.Equal(u1, u3) {
		fmt.Println("Will never happen")
	}

	if uuid.Compare(uuid.NameSpaceDNS, uuid.NameSpaceDNS) == 0 {
		fmt.Println("They are equal")
	}

	// Default Format is Canonical
	fmt.Println(uuid.Formatter(u5, uuid.CanonicalCurly))

	uuid.SwitchFormat(uuid.CanonicalBracket)

## Coverage

go test -coverprofile cover.out github.com/twinj/uuid
go tool cover -html=cover.out -o cover.html

## Links

* [RFC 4122](http://www.ietf.org/rfc/rfc4122.txt)
* [DCE 1.1: Authentication and Security Services](http://pubs.opengroup.org/onlinepubs/9629399/apdxa.htm)

## Copyright

Copyright (C) 2014 twinj@github.com
See [LICENSE](https://github.com/twinj/uuid/tree/master/LICENSE)
file for details.
