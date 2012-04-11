package goperf

import (
	"bytes"
	//"fmt"
	"testing"
)

/*
How should i read/write bytes?   byte.buffer, append slices, etc?


cat /dev/urandom | tr -dc _A-Z-a-z-0-9 | head -c800
go test -bench=".*" 
go test -bench="Byte"

*/
var (
	// this should be 800 bytes
	bydata []byte = []byte("gy-8KUvh2M_X4dV6Zj5M9Fsu7QYpQVntQ3roF0OC70_p6lIPW5Lp2zsVL3CJcPGRyWCnSMmA6XRvybRAumzuEuDI49BB8S5thTHq5xdQLqznPT0dhogg2OWB3oXA7yQ7bw1sVvstuQBlsvadnx2iJ1ZKQOpCz3tLUv-OQrwoJpC8Fi65i6W2E6qTxyr0l8a8awsnUEkL7npdbcjVTrdfVeRZkJgAwejLr27alTN16olKA5PIhCZxxbdfKwntXaqZajWP2fvWpyXZ0FLZIVtgwds7Eovw_1FoJjZmjgljgj3uNpxvdq4SGYiBY30fORzzqOz_Lw0xAdBibBtx8KNVBzJaeVQbitJIg6H2tVr14WYx3EAZkBD1kZVC9FjBgl-QA3VG1OGhrQxi5eYSPl6em1BlkCl60gyrt9BpcnGDzRzTMEzzDeGt3EdM0hs-Qfjifar3Po3Za_QdLd9lFVPOg8k_cgOSKEeeHIz4yd2pMmNbsyLw2whaMXsKvRyZd8Gg6zHMkgoK9LPcI51PFA5nUuFZLYF1-St5b2vk7GvH5wN4cTQRJVNldBVGzGjV8YPXM9tFZ8EX0wU7ZVYX8fo00c8KpNA8Q166QJtt-ksyToYZdb5W-p7_XNYpZYRRXlgEg13B_FsGQgCsL1c28HObvSBxGmENRzOl9_szLX7S5O4V4wcDWAWP2wZGmRhlLKFHVBTwGJIgS7i0swr1Pl9TdgZOyxtKk_ZBQtoMEWA5GkEF7N1cGuI86LmTCOUokLHTJgynlB1u4SZGZG5c-xdrXEBpTbbHWCSD")
)

func init() {

}

func BenchmarkByteBufWrite(b *testing.B) {
	ct := 0
	buf := bytes.Buffer{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		if ct == 10000 {
			ct = 0
			buf = bytes.Buffer{}
		}
		ct++
		buf.Write(bydata)
	}
}

// BenchmarkByteBufWrite	  500000	      2324 ns/op

func BenchmarkByteSliceAppend(b *testing.B) {
	ct := 0
	buf := make([]byte, 0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		if ct == 10000 {
			ct = 0
			buf = make([]byte, 0)
		}
		ct++
		buf = append(buf, bydata...)
	}
}

// BenchmarkByteSliceAppend	  500000	      5089 ns/op

func BenchmarkByteBufReader(b *testing.B) {
	var ct uint64
	bufx := make([]byte, 0)
	for i := 0; i < 10000; i++ {
		bufx = append(bufx, bydata...)
		bufx = append(bufx, '\n')
	}
	buf := bytes.NewBuffer(bufx)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if line, err := buf.ReadBytes('\n'); err == nil {
			ct += uint64(len(line))
		} else {
			//fmt.Println("new buffer ")
			buf = bytes.NewBuffer(bufx)
		}
	}
	//fmt.Println("total read: ", ct)
}

// BenchmarkByteBufReader	 1000000	      1586 ns/op

func BenchmarkByteSliceReader(b *testing.B) {
	var ct uint64
	pos := 0
	bufx := make([]byte, 0)
	for i := 0; i < 10000; i++ {
		bufx = append(bufx, bydata...)
		bufx = append(bufx, '\n')
	}
	size := len(bufx)
	buf := make([]byte, len(bufx))
	copy(buf, bufx)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//idx := bytes.Index(buf[pos + 1:], '\n')

		if idx := bytes.Index(buf[pos:], []byte{'\n'}); idx > -1 {
			line := buf[pos : pos+idx]
			//fmt.Println(idx, pos)
			pos += idx + 1
			if pos >= size-2 {
				buf = make([]byte, size)
				copy(buf, bufx)
				pos = 0
				//fmt.Println("starting new buf")
			}
			ct += uint64(len(line) + 1)
		}
	}
	//fmt.Println("total read: ", ct)
}

// BenchmarkByteSliceReader	1000000	      1507 ns/op
