package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

//ClientState te
type ClientState struct {
	username string
	//status string //away;busy;idle
}

var (
	newConnections  chan net.Conn
	msgChannel      chan *message
	deadConnections chan net.Conn
)

func init() {
	newConnections = make(chan net.Conn)
	msgChannel = make(chan *message)
	deadConnections = make(chan net.Conn)
}

func main() {
	//this is the count of users that ever connected.
	clientCount := 0
	var serverPort int

	flag.IntVar(&serverPort, "port", 6000, "The port to run Chat server on")
	flag.Parse()

	// Create a new Client Manager.
	cm := NewCM()
	go cm.Start()

	// Start the TCP server
	server, err := net.Listen("tcp", ":"+strconv.Itoa(serverPort))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		log.Printf("Started Server on port %d...", serverPort)
	}
	defer func() {
		server.Close()
		log.Println("Server/Listener closed")
	}()
	// Tell the server to accept connections forever
	// and push new connections into the newConnections channel.
	//
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			clientCount++
			// Add this connection to the `clientsMap` map
			//
			//clientsMap[conn] =
			if cm.Write(conn, "anonymous"+strconv.Itoa(clientCount)) {
				log.Printf("Accepted new client, %s@%s", "anonymous"+strconv.Itoa(clientCount), conn.RemoteAddr())
				sendWelcome(conn)
				newConnections <- conn
			}
		}
	}()

	for {
		// Handle 1) new connections; 2) dead connections;
		// and, 3) received messages.
		//
		select {
		case conn := <-newConnections:
			cs := cm.Read(conn)
			go cm.handleMessages(conn, cs)

		case msg := <-msgChannel:
			cm.msgsForBroadcast <- &broadcastMsgOp{msg: msg}

		case conn := <-deadConnections:
			cm.kills <- conn
		}

	}
}
