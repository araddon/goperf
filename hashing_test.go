package goperf

import (
	"crypto/md5"
	"github.com/dchest/siphash"
	"testing"
)

func BenchmarkMd5Hash(b *testing.B) {
	h := md5.New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h.Write([]byte("hello world"))
		h.Sum(nil)
		h.Reset()
	}
}

// 1.02 BenchmarkMd5Hash	 1000000	      1330 ns/op
// 1.10 BenchmarkMd5Hash	 5000000	       625 ns/op

func BenchmarkSipHash(b *testing.B) {
	h := siphash.New([]byte("0123456789012345"))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h.Write([]byte("Hello world"))
		h.Sum(nil)
		h.Reset()
	}
}

// 1.02 BenchmarkSipHash	 5000000	       478 ns/op
// 1.10 BenchmarkSipHash	10000000	       251 ns/op

func BenchmarkSipHashFast(b *testing.B) {
	var key0, key1 uint64 = 0, 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		siphash.Hash(key0, key1, []byte("Hello World"))
	}
}

// 1.02 BenchmarkSipHashFast	20000000	        73.5 ns/op
// 1.10 BenchmarkSipHashFast	50000000	        38.3 ns/op
