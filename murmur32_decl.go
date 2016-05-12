// +build go1.5,amd64

package murmur3

func Sum32(data []byte) (h1 uint32)

func SeedSum32(seed uint32, data []byte) (h1 uint32)
