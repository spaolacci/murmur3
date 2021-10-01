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

type digest struct {
	clen int      // Digested input cumulative length.
	tail []byte   // 0 to Size()-1 bytes view of `buf'.
	buf  [16]byte // Expected (but not required) to be Size() large.
	seed uint32   // Seed for initializing the hash.
}

func (d *digest) BlockSize() int { return 1 }

func (d *digest) write(p []byte, size int, bmix func([]byte) []byte) (n int, err error) {
	n = len(p)
	d.clen += n

	if len(d.tail) > 0 {
		// Stick back pending bytes.
		nfree := size - len(d.tail) // nfree ∈ [1, d.Size()-1].
		if nfree < len(p) {
			// One full block can be formed.
			block := append(d.tail, p[:nfree]...)
			p = p[nfree:]
			_ = bmix(block) // No tail.
		} else {
			// Tail's buf is large enough to prevent reallocs.
			p = append(d.tail, p...)
		}
	}

	d.tail = bmix(p)

	// Keep own copy of the 0 to Size()-1 pending bytes.
	nn := copy(d.buf[:], d.tail)
	d.tail = d.buf[:nn]

	return n, nil
}

func (d *digest) reset() {
	d.clen = 0
	d.tail = nil
}
