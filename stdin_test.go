package goperf

import (
	//"bytes"
	"log"
	"os/exec"
	"testing"
	//"time"
	"bufio"
)

//go test -bench="Stdio"

func BenchmarkStdioPy(b *testing.B) {
	cmd := exec.Command("python", "msg.py", "hello")

	out, err := cmd.StdoutPipe()
	rdr := bufio.NewReader(out)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rdr.ReadString('\n')
	}
}

// BenchmarkStdioPy	    1000000	    1641 ns/op
//  = 609,385/sec

func BenchmarkStdioGo(b *testing.B) {
	cmd := exec.Command("go", "run", "cmds/stdio.go")

	out, err := cmd.StdoutPipe()
	rdr := bufio.NewReader(out)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rdr.ReadString('\n')
	}
}

// BenchmarkStdioGo	    1000000	      1051 ns/op
//  = 951,475/sec
