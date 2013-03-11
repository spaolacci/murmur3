murmur3
=======

Native Go implementation of Austin Appleby's third MurmurHash revision (aka
MurmurHash3).

Reference algorithm has been slightly hacked as to support the streaming mode
required by Go's standard [Hash interface](http://golang.org/pkg/hash/#Hash).


Benchmarks
----------

Go tip as of 2013-03-11 (i.e almost go1.1), core i7 @ 3.4 Ghz. All runs
include hasher instanciation and sequence finalization.

<pre>

Benchmark32_1        50000000        26.7 ns/op        37.47 MB/s
Benchmark32_2       100000000        28.2 ns/op        70.95 MB/s
Benchmark32_4       100000000        26.4 ns/op       151.56 MB/s
Benchmark32_8        50000000        29.0 ns/op       275.42 MB/s
Benchmark32_16       50000000        32.3 ns/op       495.74 MB/s
Benchmark32_32       50000000        39.1 ns/op       818.67 MB/s
Benchmark32_64       50000000        54.1 ns/op      1182.40 MB/s
Benchmark32_128      20000000        84.3 ns/op      1518.96 MB/s
Benchmark32_256      10000000       150.0 ns/op      1700.58 MB/s
Benchmark32_512      10000000       272.0 ns/op      1877.52 MB/s
Benchmark32_1024      5000000       520.0 ns/op      1968.68 MB/s
Benchmark32_2048      1000000      1012.0 ns/op      2022.71 MB/s
Benchmark32_4096      1000000      1998.0 ns/op      2049.31 MB/s
Benchmark32_8192       500000      3973.0 ns/op      2061.41 MB/s

Benchmark128_1       50000000        33.2 ns/op        30.15 MB/s
Benchmark128_2       50000000        33.3 ns/op        59.99 MB/s
Benchmark128_4       50000000        35.4 ns/op       112.88 MB/s
Benchmark128_8       50000000        36.6 ns/op       218.30 MB/s
Benchmark128_16      50000000        35.5 ns/op       450.86 MB/s
Benchmark128_32      50000000        35.3 ns/op       905.84 MB/s
Benchmark128_64      50000000        44.3 ns/op      1443.76 MB/s
Benchmark128_128     50000000        58.2 ns/op      2201.02 MB/s
Benchmark128_256     20000000        85.3 ns/op      2999.88 MB/s
Benchmark128_512     10000000       142.0 ns/op      3592.97 MB/s
Benchmark128_1024    10000000       258.0 ns/op      3963.74 MB/s
Benchmark128_2048     5000000       494.0 ns/op      4144.65 MB/s
Benchmark128_4096     2000000       955.0 ns/op      4285.80 MB/s
Benchmark128_8192     1000000      1884.0 ns/op      4347.12 MB/s

</pre>
