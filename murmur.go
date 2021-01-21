// Copyright 2013, Sébastien Paolacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package murmur3 implements Austin Appleby's non-cryptographic MurmurHash3.

 Reference implementation:
    http://code.google.com/p/smhasher/wiki/MurmurHash3

 History, characteristics and (legacy) perfs:
    https://sites.google.com/site/murmurhash/
    https://sites.google.com/site/murmurhash/statistics
*/
package murmur3

import (
	"unsafe"
)

type bmixer interface {
	bmix(p []byte) (tail []byte)
	Size() (n int)
	reset()
}

type digest struct {
	clen int      // Digested input cumulative length.
	tail []byte   // 0 to Size()-1 bytes view of `buf'.
	buf  [16]byte // Expected (but not required) to be Size() large.
	seed uint32   // Seed for initializing the hash.
	bmixer
}

// sliceHeader is similar to reflect.SliceHeader, but it assumes that the layout
// of the first two words is the same as the layout of a string.
type sliceHeader struct {
	s   string
	cap int
}

func (d *digest) BlockSize() int { return 1 }

func (d *digest) WriteString(s string) (int, error) {
	// This code does the same as:
	//
	//   return d.Write([]byte(s))
	//
	// because the parameter `p` to Write is passed to d.bmix, which is an
	// interface type, the simple conversion to a byte slice forces the compiler
	// to make a heap allocation and copy the slice. The use of unsafe here lets
	// us take over the default behavior of the compiler to have the byte slice
	// share the underlying memory buffer of the string and avoid the extra heap
	// allocation.
	return d.Write(*(*[]byte)(unsafe.Pointer(&sliceHeader{s: s, cap: len(s)})))
}

func (d *digest) Write(p []byte) (n int, err error) {
	n = len(p)
	d.clen += n

	if len(d.tail) > 0 {
		// Stick back pending bytes.
		nfree := d.Size() - len(d.tail) // nfree ∈ [1, d.Size()-1].
		if nfree < len(p) {
			// One full block can be formed.
			block := append(d.tail, p[:nfree]...)
			p = p[nfree:]
			_ = d.bmix(block) // No tail.
		} else {
			// Tail's buf is large enough to prevent reallocs.
			p = append(d.tail, p...)
		}
	}

	d.tail = d.bmix(p)

	// Keep own copy of the 0 to Size()-1 pending bytes.
	nn := copy(d.buf[:], d.tail)
	d.tail = d.buf[:nn]

	return n, nil
}

func (d *digest) Reset() {
	d.clen = 0
	d.tail = nil
	d.bmixer.reset()
}
