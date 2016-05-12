murmur3
=======

Native Go implementation of Austin Appleby's third MurmurHash revision (aka
MurmurHash3).

Includes assembly for amd64 (go 1.5+), the benchmarks of which can be seen in
PR [#1](https://github.com/twmb/murmur3/pull/1).

Reference algorithm has been slightly hacked as to support the streaming mode
required by Go's standard [Hash interface](http://golang.org/pkg/hash/#Hash).

Documentation can be found on [godoc](https://godoc.org/github.com/twmb/murmur3).
