package main

import (
	"flag"
)

func main() {
	var serverPort int

	flag.IntVar(&serverPort, "port", 6000, "The port to run the Chat server on")
	flag.Parse()

	server, err := NewChatServer(serverPort)
	if err != nil {
		panic(err.Error())
	}
	server.Start()
}
