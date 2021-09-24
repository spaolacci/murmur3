package murmur3

// http://code.google.com/p/guava-libraries/source/browse/guava/src/com/google/common/hash/Murmur3_32HashFunction.java

import (
	"hash"
	"math/bits"
	"unsafe"
)

// Make sure interfaces are correctly implemented.
var (
	_ hash.Hash   = new(Digest32)
	_ hash.Hash32 = new(Digest32)
	_ bmixer      = new(Digest32)
)

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593
)

// Digest32 represents a partial evaluation of a 32 bites hash.
type Digest32 struct {
	digest
	h1 uint32 // Unfinalized running hash.
}

// New32 returns new 32-bit hasher
func New32() *Digest32 { return New32WithSeed(0) }

// New32WithSeed returns new 32-bit hasher set with explicit seed value
func New32WithSeed(seed uint32) *Digest32 {
	d := new(Digest32)
	d.seed = seed
	d.bmixer = d
	d.Reset()
	return d
}

func (d *Digest32) Size() int { return 4 }

func (d *Digest32) reset() { d.h1 = d.seed }

func (d *Digest32) Sum(b []byte) []byte {
	h := d.Sum32()
	return append(b, byte(h>>24), byte(h>>16), byte(h>>8), byte(h))
}

// digest as many blocks as possible.
func (d *Digest32) bmix(p []byte) (tail []byte) {
	h1 := d.h1

	nblocks := len(p) / 4
	for i := 0; i < nblocks; i++ {
		k1 := *(*uint32)(unsafe.Pointer(&p[i*4]))

		k1 *= c1_32
		k1 = bits.RotateLeft32(k1, 15)
		k1 *= c2_32

		h1 ^= k1
		h1 = bits.RotateLeft32(h1, 13)
		h1 = h1*4 + h1 + 0xe6546b64
	}
	d.h1 = h1
	return p[nblocks*d.Size():]
}

func (d *Digest32) Sum32() (h1 uint32) {

	h1 = d.h1

	var k1 uint32
	switch len(d.tail) & 3 {
	case 3:
		k1 ^= uint32(d.tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(d.tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(d.tail[0])
		k1 *= c1_32
		k1 = bits.RotateLeft32(k1, 15)
		k1 *= c2_32
		h1 ^= k1
	}

	h1 ^= uint32(d.clen)

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}

// Sum32 returns the MurmurHash3 sum of data. It is equivalent to the
// following sequence (without the extra burden and the extra allocation):
//     hasher := New32()
//     hasher.Write(data)
//     return hasher.Sum32()
func Sum32(data []byte) uint32 { return Sum32WithSeed(data, 0) }

// Sum32WithSeed returns the MurmurHash3 sum of data. It is equivalent to the
// following sequence (without the extra burden and the extra allocation):
//     hasher := New32WithSeed(seed)
//     hasher.Write(data)
//     return hasher.Sum32()
func Sum32WithSeed(data []byte, seed uint32) uint32 {

	h1 := seed

	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}
	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		k1 := *(*uint32)(unsafe.Pointer(p))

		k1 *= c1_32
		k1 = bits.RotateLeft32(k1, 15)
		k1 *= c2_32

		h1 ^= k1
		h1 = bits.RotateLeft32(h1, 13)
		h1 = h1*4 + h1 + 0xe6546b64
	}

	tail := data[nblocks*4:]

	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1_32
		k1 = bits.RotateLeft32(k1, 15)
		k1 *= c2_32
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}
