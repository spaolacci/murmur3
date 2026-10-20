package murmur3_test

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

func ExampleSum128WithSeed() {
	h1, h2 := murmur3.Sum128WithSeed([]byte("hello"), 42)
	fmt.Printf("%016x%016x\n", h1, h2)

	h1, h2 = murmur3.Sum128WithSeed(nil, 0)
	fmt.Printf("%016x%016x\n", h1, h2)
	// Output:
	// c4b8b3c960af6f082334b875b0efbc7a
	// 00000000000000000000000000000000
}

func ExampleNew128WithSeed() {
	h := murmur3.New128WithSeed(42)
	h.Write([]byte("hel"))
	h.Write([]byte("lo"))
	fmt.Printf("%x\n", h.Sum(nil))
	// Output: c4b8b3c960af6f082334b875b0efbc7a
}
