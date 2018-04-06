package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type clientState struct {
	username string
	//status string //away;busy;idle
}

var (
	//client state -
	/*
		reads            chan *readOp
		readsAllVals     chan *readOpAllVals
		writes           chan *writeOp
		msgsForBroadcast chan *broadcastMsgOp
		kills            chan net.Conn
	*/
	//
	newConnections  chan net.Conn
	msgChannel      chan *message
	deadConnections chan net.Conn
)

func init() {
	/*
		//client state
		reads = make(chan *readOp)
		readsAllVals = make(chan *readOpAllVals)
		writes = make(chan *writeOp)
		msgsForBroadcast = make(chan *broadcastMsgOp)
		kills = make(chan net.Conn)
		//
	*/

	newConnections = make(chan net.Conn)
	msgChannel = make(chan *message)
	deadConnections = make(chan net.Conn)
}

func main() {
	//this is the count of users that ever connected.
	clientCount := 0
	serverPort := 6000

	cliManager := clientManager{
		clients:          make(map[net.Conn]*clientState),
		reads:            make(chan *readOp),
		readsAllVals:     make(chan *readOpAllVals),
		writes:           make(chan *writeOp),
		msgsForBroadcast: make(chan *broadcastMsgOp),
		kills:            make(chan net.Conn),
	}
	go cliManager.start()
	// Start the TCP server
	//
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
			if cliManager.submitCMWrite(conn, "anonymous"+strconv.Itoa(clientCount)) {
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
			cs := cliManager.submitCMRead(conn)
			go cliManager.handleMessages(conn, cs)

		case msg := <-msgChannel:
			cliManager.msgsForBroadcast <- &broadcastMsgOp{msg: msg}

		case conn := <-deadConnections:
			cliManager.kills <- conn
		}

	}
}
