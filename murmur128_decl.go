// +build amd64

package murmur3

func Sum128(data []byte) (h1 uint64, h2 uint64)

func SeedSum128(seed1, seed2 uint64, data []byte) (h1 uint64, h2 uint64)
