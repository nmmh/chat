package main

import (
	"fmt"
	"log"
	"net"
)

type readOp struct {
	key  net.Conn
	resp chan *clientState
}

type readOpAllVals struct {
	resp chan []string
}
type writeOp struct {
	key  net.Conn
	val  *clientState
	resp chan bool
}

type broadcastMsgOp struct {
	msg *message
}
type clientManager struct {
	clients          map[net.Conn]*clientState
	reads            chan *readOp
	readsAllVals     chan *readOpAllVals
	writes           chan *writeOp
	msgsForBroadcast chan *broadcastMsgOp
	kills            chan net.Conn
}

func (manager *clientManager) start() {
	for {
		select {
		case read := <-manager.reads:
			read.resp <- manager.clients[read.key]
		case readAllVals := <-manager.readsAllVals:
			s := make([]string, 0)
			for _, val := range manager.clients {
				s = append(s, val.username)
			}
			readAllVals.resp <- s
		case write := <-manager.writes:
			manager.clients[write.key] = write.val
			write.resp <- true
		case msgForBroadcast := <-manager.msgsForBroadcast:
			// Loop over all connected clients
			for conn := range manager.clients {
				if msgForBroadcast.msg.msgScope == "ALLEXCEPTSENDER" {
					if msgForBroadcast.msg.username == manager.clients[conn].username {
						continue
					}
				} else if msgForBroadcast.msg.msgScope == "SENDERONLY" {
					if msgForBroadcast.msg.username != manager.clients[conn].username {
						continue
					}
				}
				//send msg in  a goroutine
				go sendMessage(conn, msgForBroadcast.msg)
			}
			//message always logged at the server
			log.Printf("%s", msgForBroadcast.msg.text)
		case kill := <-manager.kills:
			msgChannel <- &message{manager.clients[kill].username, "CHANOP", "ALL", fmt.Sprintf(" * [%s] disconnected\r\n", manager.clients[kill].username)}
			delete(manager.clients, kill)
			//log.Printf(getUserList(clientsMap))
			kill.Close()
		}
	}
}

func (manager *clientManager) submitCMRead(conn net.Conn) *clientState {
	read := &readOp{key: conn, resp: make(chan *clientState)}
	manager.reads <- read
	return <-read.resp
}

func (manager *clientManager) submitCMReadAll() []string {
	readAllVals := &readOpAllVals{resp: make(chan []string)}
	manager.readsAllVals <- readAllVals
	return <-readAllVals.resp
}

func (manager *clientManager) submitCMWrite(conn net.Conn, username string) bool {
	write := &writeOp{key: conn, val: &clientState{username: username}, resp: make(chan bool)}
	manager.writes <- write
	return <-write.resp
}
