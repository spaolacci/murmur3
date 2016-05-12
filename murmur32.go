package murmur3

import (
	"hash"
	"unsafe"
)

// Make sure interfaces are correctly implemented.
var (
	_ hash.Hash   = new(digest32)
	_ hash.Hash32 = new(digest32)
)

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593
)

// digest32 represents a partial evaluation of a 32 bites hash.
type digest32 struct {
	digest
	h1 uint32 // Unfinalized running hash.
}

// SeedNew32 returns a hash.Hash32 for streaming 32 bit sums with its internal
// digest initialized to seed.
func SeedNew32(seed uint32) hash.Hash32 {
	d := &digest32{h1: seed}
	d.bmixer = d
	d.Reset()
	return d
}

// New32 returns a hash.Hash32 for streaming 32 bit sums.
func New32() hash.Hash32 {
	return SeedNew32(0)
}

func (d *digest32) Size() int { return 4 }

func (d *digest32) reset() { d.h1 = 0 }

func (d *digest32) Sum(b []byte) []byte {
	h := d.Sum32()
	return append(b, byte(h>>24), byte(h>>16), byte(h>>8), byte(h))
}

// Digest as many blocks as possible.
func (d *digest32) bmix(p []byte) (tail []byte) {
	h1 := d.h1

	nblocks := len(p) / 4
	for i := 0; i < nblocks; i++ {
		k1 := *(*uint32)(unsafe.Pointer(&p[i*4]))

		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19) // rotl32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}
	d.h1 = h1
	return p[nblocks*d.Size():]
}

func (d *digest32) Sum32() (h1 uint32) {

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
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
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
