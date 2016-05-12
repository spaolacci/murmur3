// +build go1.5,amd64

package murmur3

// Sum128 returns the MurmurHash3 sum of data. It is equivalent to the
// following sequence (without the extra burden and the extra allocation):
//     hasher := New128()
//     hasher.Write(data)
//     return hasher.Sum128()
func Sum128(data []byte) (h1 uint64, h2 uint64)

// SeedSum128 returns the MurmurHash3 sum of data with digests initialized to
// seed1 and seed2.
func SeedSum128(seed1, seed2 uint64, data []byte) (h1 uint64, h2 uint64)
