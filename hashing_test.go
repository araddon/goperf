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

//BenchmarkMd5Hash	 1000000	      1330 ns/op

func BenchmarkSipHash(b *testing.B) {
	h := siphash.New([]byte("0123456789012345"))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h.Write([]byte("Hello world"))
		h.Sum(nil)
		h.Reset()
	}
}

//BenchmarkSipHash	 5000000	       478 ns/op

func BenchmarkSipHashFast(b *testing.B) {
	var key0, key1 uint64 = 0, 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		siphash.Hash(key0, key1, []byte("Hello World"))
	}
}

//BenchmarkSipHashFast	20000000	        73.5 ns/op
