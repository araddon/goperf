package main

import (
	zmq "github.com/alecthomas/gozmq"
)

func main() {

	context, _ := zmq.NewContext()
	defer context.Close()

	socketPub, _ := context.NewSocket(zmq.PUB)
	defer socketPub.Close()
	socketPub.Bind("tcp://*:7567")
	println("starting zmq publisher on tcp://*:7567")

	for {
		socketPub.Send([]byte("hello"), 0)
	}
}
