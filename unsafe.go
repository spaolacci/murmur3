package murmur3

import "unsafe"

func unsafeStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&sliceHeader{str: s, cap: len(s)}))
}

type sliceHeader struct {
	str string
	cap int
}
