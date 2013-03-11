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

Benchmark32_1      20000000     118 ns/op           8.46 MB/s 
Benchmark32_2      20000000     134 ns/op          14.88 MB/s 
Benchmark32_4      20000000     116 ns/op          34.21 MB/s 
Benchmark32_8      20000000     118 ns/op          67.65 MB/s 
Benchmark32_16     20000000     121 ns/op         131.31 MB/s 
Benchmark32_32     20000000     129 ns/op         246.21 MB/s 
Benchmark32_64     10000000     145 ns/op         440.35 MB/s 
Benchmark32_128    10000000     177 ns/op         723.01 MB/s 
Benchmark32_256    10000000     246 ns/op        1037.15 MB/s 
Benchmark32_512     5000000     375 ns/op        1363.39 MB/s 
Benchmark32_1024    5000000     627 ns/op        1631.87 MB/s 
Benchmark32_2048    1000000    1105 ns/op        1853.38 MB/s 
Benchmark32_4096    1000000    2091 ns/op        1958.09 MB/s 
Benchmark32_8192     500000    4062 ns/op        2016.45 MB/s 

Benchmark128_1     20000000     122 ns/op           8.14 MB/s 
Benchmark128_2     10000000     138 ns/op          14.48 MB/s 
Benchmark128_4     10000000     159 ns/op          25.01 MB/s 
Benchmark128_8     10000000     143 ns/op          55.92 MB/s 
Benchmark128_16    20000000     122 ns/op         131.09 MB/s 
Benchmark128_32    20000000     125 ns/op         254.00 MB/s 
Benchmark128_64    20000000     133 ns/op         479.22 MB/s 
Benchmark128_128   10000000     149 ns/op         856.81 MB/s 
Benchmark128_256   10000000     181 ns/op        1408.09 MB/s 
Benchmark128_512   10000000     240 ns/op        2132.96 MB/s 
Benchmark128_1024   5000000     360 ns/op        2843.22 MB/s 
Benchmark128_2048   5000000     606 ns/op        3374.67 MB/s 
Benchmark128_4096   1000000    1064 ns/op        3849.06 MB/s 
Benchmark128_8192   1000000    1994 ns/op        4106.64 MB/s 

</pre>
