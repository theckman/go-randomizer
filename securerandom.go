// Copyright 2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

// Package securerandom is a set of utilities for generating random data from the
// secure sources provided by the "crypto/rand" package. The package has the
// ability to generate bytes from securerandom data, as well as other data from
// those bytes. That includes numbers and Base64-encoded strings.
//
// Example:
//
//		import (
//			"math/rand"
//			"github.com/theckman/go-securerandom"
//		)
//
// 		ri64, err := securerandom.Int64()
//
// 		// secure-random data is unavailable
// 		if err != nil {
// 			// handle err
// 		}
//		rand.Seed(ri64)
package securerandom

import (
	crand "crypto/rand"
	"encoding/base64"
	mrand "math/rand"
)

// Bytes is a function that takes an integer and returns
// a slice of that length containing random bytes.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)

	if _, err := crand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

// maximumBytes is used to calculate the how many bytes we can generate a
// base64 string from, if we don't want the string to be longer than size.
func maximumBytes(size int) int {
	return int((float64(size) / 4) * 3)
}

// Base64 is a function that returns a randomize standard Base64
// string. This does not use the URL encoding. It takes a single parameter that
// is the maximum possible length of the string, it will get as close as possible.
func Base64(max int) (string, error) {
	b, err := Bytes(maximumBytes(max))
	return base64.StdEncoding.EncodeToString(b), err
}

// URLBase64 is a function that returns a random URL encoded Base64
// string. This does not use the URL encoding. It takes a single parameter that
// is the maximum possible length of the string, it will get as close as possible.
func URLBase64(max int) (string, error) {
	b, err := Bytes(maximumBytes(max))
	return base64.URLEncoding.EncodeToString(b), err
}

// Uint16 is a function that returns a uint16 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func Uint16() (uint16, error) {
	b, err := Bytes(2)

	if err != nil {
		return 0, err
	}

	var u16 uint16

	for i := range b {
		offset := uint16(i) + 1
		shift := 16 - (8 * offset)
		u16 = u16 | uint16(b[i])<<shift
	}

	return u16, nil
}

// Uint32 is a function that returns a uint32 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func Uint32() (uint32, error) {
	b, err := Bytes(4)

	if err != nil {
		return 0, err
	}

	var u32 uint32

	for i := range b {
		offset := uint32(i) + 1
		shift := 32 - (8 * offset)
		u32 = u32 | uint32(b[i])<<shift
	}

	return u32, nil
}

// Uint64 is a function that returns a uint64 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func Uint64() (uint64, error) {
	b, err := Bytes(8)

	if err != nil {
		return 0, err
	}

	var u64 uint64

	for i := range b {
		offset := uint64(i) + 1
		shift := 64 - (8 * offset)
		u64 = u64 | uint64(b[i])<<shift
	}

	return u64, nil
}

// Int16 is a function that returns a int16 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func Int16() (int16, error) {
	b, err := Bytes(2)

	if err != nil {
		return 0, err
	}

	var i16 int16

	for i := range b {
		offset := uint16(i) + 1
		shift := 16 - (8 * offset)
		i16 = i16 | int16(b[i])<<shift
	}

	return i16, nil
}

// Int32 is a function that returns a int32 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func Int32() (int32, error) {
	b, err := Bytes(4)

	if err != nil {
		return 0, err
	}

	var i32 int32

	for i := range b {
		offset := uint32(i) + 1
		shift := 32 - (8 * offset)
		i32 = i32 | int32(b[i])<<shift
	}

	return i32, nil
}

// Int64 is a function that returns a int64 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func Int64() (int64, error) {
	b, err := Bytes(8)

	if err != nil {
		return 0, err
	}

	var i64 int64

	for i := range b {
		offset := uint64(i) + 1
		shift := 64 - (8 * offset)
		i64 = i64 | int64(b[i])<<shift
	}

	return i64, nil
}

// RandSource is a function that returns a Source from the "math/rand" package
// to be used to create a new pseudorandom generator. If this returns err != nil
// the value of the source is not suitable for use.
func RandSource() (mrand.Source, error) {
	randInt64, err := Int64()

	if err != nil {
		return nil, err
	}

	return mrand.NewSource(randInt64), nil
}
