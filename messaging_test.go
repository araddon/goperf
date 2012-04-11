/*
To run this test, go to zmqio and go run main.go 
Then (seperate console) run this test
*/
package goperf

import (
	"testing"

	zmq "github.com/alecthomas/gozmq"
)

//go test -bench="InProc"

/*
	"log"
	"os/exec"
	"time"
var (
	zmqInProcReceiver func(string)
	zmqInProcSender   func(string)
	inprocCt          int
)

func init() {
	zmqInProcReceiver = func(msg string) {
		inprocCt++
	}
	zmqInProcSender = RunZmq("inproc://gozmq", "inproc://gozmq", zmqInProcReceiver)
	//cmd := exec.Command("go", "run", "main.go")
	cmd := exec.Command("zmqio")
	go func() {
		out, err := cmd.Output()
		log.Println(string(out), err)
	}()
}

func BenchmarkZmqInProc(b *testing.B) {
	inprocCt = 0

	go func() {
		for i := 0; i < 1e9; i++ {
			zmqInProcSender("hello")
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if i > inprocCt {
			time.Sleep(time.Millisecond * 1)
		}
	}
}

*/

//BenchmarkZmqInProc	   10000	    176783 ns/op

func RunZmq(pub, sub string, handler func(string)) func(string) {

	context, _ := zmq.NewContext()
	defer context.Close()

	socketPub, _ := context.NewSocket(zmq.PUB)
	defer socketPub.Close()
	socketPub.Bind(pub)
	println("starting zmq publisher on ", pub)

	socketSub, _ := context.NewSocket(zmq.SUB)
	defer socketSub.Close()
	// connect with no subscription filter
	socketSub.Connect(sub)
	socketSub.SetSockOptString(zmq.SUBSCRIBE, "")

	println("starting zmq subscriber on ", sub)

	go func() {
		for {
			// block here, waiting for inbound requests
			msg, _ := socketSub.Recv(0)
			handler(string(msg))
		}
	}()

	return func(msg string) {
		socketPub.Send([]byte(msg), 0)
	}
}

func BenchmarkZmqSub(b *testing.B) {
	context, _ := zmq.NewContext()
	defer context.Close()

	socketSub, _ := context.NewSocket(zmq.SUB)
	defer socketSub.Close()
	// connect with no subscription filter
	socketSub.Connect("tcp://localhost:7567")
	socketSub.SetSockOptString(zmq.SUBSCRIBE, "")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = socketSub.Recv(0)
	}
}
