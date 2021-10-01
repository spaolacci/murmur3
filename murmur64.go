package murmur3

import (
	"hash"
	"io"
)

// Make sure interfaces are correctly implemented.
var (
	_ hash.Hash       = (*Digest64)(nil)
	_ hash.Hash64     = (*Digest64)(nil)
	_ io.StringWriter = (*Digest64)(nil)
)

// Digest64 is half a Digest128.
type Digest64 Digest128

// New64 returns a 64-bit hasher
func New64() *Digest64 { return New64WithSeed(0) }

// New64WithSeed returns a 64-bit hasher set with explicit seed value
func New64WithSeed(seed uint32) *Digest64 {
	d := (*Digest64)(New128WithSeed(seed))
	return d
}

func (d *Digest64) Sum(b []byte) []byte {
	h1 := d.Sum64()
	return append(b,
		byte(h1>>56), byte(h1>>48), byte(h1>>40), byte(h1>>32),
		byte(h1>>24), byte(h1>>16), byte(h1>>8), byte(h1))
}

func (d *Digest64) Sum64() uint64 {
	h1, _ := (*Digest128)(d).Sum128()
	return h1
}

func (d *Digest64) Size() int { return 8 }

func (d *Digest64) WriteString(s string) (int, error) {
	return d.Write(unsafeStringToBytes(s))
}

func (d *Digest64) Write(b []byte) (int, error) {
	return d.write(b, 16, (*Digest128)(d).bmix)
}

func (d *Digest64) Reset() {
	d.reset()
	d.h1, d.h2 = uint64(d.seed), uint64(d.seed)
}

// Sum64 returns the MurmurHash3 sum of data. It is equivalent to the
// following sequence (without the extra burden and the extra allocation):
//     hasher := New64()
//     hasher.Write(data)
//     return hasher.Sum64()
func Sum64(data []byte) uint64 { return Sum64WithSeed(data, 0) }

// Sum64WithSeed returns the MurmurHash3 sum of data. It is equivalent to the
// following sequence (without the extra burden and the extra allocation):
//     hasher := New64WithSeed(seed)
//     hasher.Write(data)
//     return hasher.Sum64()
func Sum64WithSeed(data []byte, seed uint32) uint64 {
	d := Digest128{h1: uint64(seed), h2: uint64(seed)}
	d.seed = seed
	d.tail = d.bmix(data)
	d.clen = len(data)
	h1, _ := d.Sum128()
	return h1
}
