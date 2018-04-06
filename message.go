package main

import (
	"io/ioutil"
	"net"
	"path/filepath"
)

type message struct {
	username string
	msgType  string //NORMAL;CHANOP;WHSIPER
	msgScope string //ALL;SENDERONLY,ALLEXCEPTSENDER;P2P
	text     string
}

func getWelcome() []byte {
	fileName, _ := filepath.Abs("banner.txt")
	//fileName := "C:/neil/dev/gowork/src/chat/banner.txt"
	welcome, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return append(welcome, "\r\n"...)
}

func sendWelcome(conn net.Conn) {

	//send welcome message in a goroutine so that the network doesnt block
	go sendString(conn, string(getWelcome()))
}
func sendMessage(conn net.Conn, msg *message) {
	sendString(conn, msg.text)
}

func sendString(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		//deadConnections <- conn
	}
}
