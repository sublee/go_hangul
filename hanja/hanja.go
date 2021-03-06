// Copyright 2013, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hanja

// Convert hanjas in hangul string to hangul
func Convert(h string) string {
	rh := []rune(h)

	for i, c := range rh {
		h, ok := table[c]
		if ok {
			rh[i] = h
		}
	}

	return string(rh)
}

// Return true if can convert given hanja rune to hangul.
func IsHanja(c rune) bool {
	_, ok := table[c]
	return ok
}
