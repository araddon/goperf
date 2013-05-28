package goperf

import (
	"testing"
	//"time"
)

//  go test -bench=".*"
//  go test -bench="Channel"

func BenchmarkChannelSend(b *testing.B) {
	ch := make(chan string)
	ct := 0
	go func() {
		for _ = range ch {
			ct++
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ch <- "hello"
	}
}

/*
1.02 BenchmarkChannelSend	 5000000	       309 ns/op
1.10 BenchmarkChannelSend	10000000	       147 ns/op
*/

// No longer stable in 1.1
// func BenchmarkGoRoutineSend(b *testing.B) {
// 	ct := 0
// 	doit := func(msg string) {
// 		ct++
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		go doit("hello")
// 	}
// }

/*
BenchmarkGoRoutineSend	  500000	      3491 ns/op
*/
// no longer stable in 1.1
// func BenchmarkGoRoutineMaxct(b *testing.B) {
// 	ct := 0
// 	doit := func(msg string) {
// 		ct++
// 		time.Sleep(time.Second * 1)
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		go doit("hello")
// 	}
// }

/*
BenchmarkGoRoutineMaxct	 1000000	     83946 ns/op
*/
