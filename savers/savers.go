// This package provides savers for the guthub.com/twinj/uuid RFC4122 UUIDs.
//
// Use this package for V1 and V2 UUIDs or your own implementation.
//
// By applying a saver you can store any UUID generation required data in a
// non volatile store, the purpose of which is to save the Clock Sequence,
// last Timestamp and the last Node id used in the last generated UUID.
//
// The Saver Save method is called every time you generate that UUID.
//
// The example code in the specification was also used as reference
// for design.
//
// Copyright (C) 2016 twinj@github.com  2014 MIT licence
package savers
