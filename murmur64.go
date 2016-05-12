package murmur3

import (
	"hash"
)

// Make sure interfaces are correctly implemented.
var (
	_ hash.Hash   = new(digest64)
	_ hash.Hash64 = new(digest64)
	_ bmixer      = new(digest64)
)

// digest64 is half a digest128.
type digest64 digest128

// SeedNew64 returns a hash.Hash64 for streaming 64 bit sums with its internal
// digest initialized to seed.
func SeedNew64(seed uint64) hash.Hash64 {
	return (*digest64)(SeedNew128(seed, 0).(*digest128))
}

// New64 returns a hash.Hash64 for streaming 64 bit sums.
func New64() hash.Hash64 {
	return SeedNew64(0)
}

func (d *digest64) Sum(b []byte) []byte {
	h1 := d.Sum64()
	return append(b,
		byte(h1>>56), byte(h1>>48), byte(h1>>40), byte(h1>>32),
		byte(h1>>24), byte(h1>>16), byte(h1>>8), byte(h1))
}

func (d *digest64) Sum64() uint64 {
	h1, _ := (*digest128)(d).Sum128()
	return h1
}

// Sum64 returns the MurmurHash3 sum of data. It is equivalent to the
// following sequence (without the extra burden and the extra allocation):
//     hasher := New64()
//     hasher.Write(data)
//     return hasher.Sum64()
func Sum64(data []byte) uint64 {
	h1, _ := SeedSum128(0, 0, data)
	return h1
}

// SeedSum64 returns the MurmurHash3 sum of data with the digest initialized to
// seed.
func SeedSum64(seed uint64, data []byte) uint64 {
	h1, _ := SeedSum128(seed, 0, data)
	return h1
}
