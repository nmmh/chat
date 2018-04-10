package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	//var serverPort int
	var servAddr string
	strEcho := "-"
	flag.StringVar(&servAddr, "addr", "127.0.0.1:6000", "The address of the server to connect to in the format \"address:port\"")
	//flag.IntVar(&serverPort, "port", 6000, "The port of the server to connect to")
	flag.Parse()
	fmt.Printf("Connecting to: %s\n", servAddr)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	time.Sleep(2 * time.Second)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	var run = true

	for i := 0; run; i++ {
		time.Sleep(200 * time.Millisecond)
		select {
		case qc := <-quitChan:
			run = false
			fmt.Printf("\nReceived an interrupt [%s], stopping services\n", qc.String())
			break
		default:
			go func(conn net.Conn, i int) {
				reader := bufio.NewReader(conn)
				for true {
					message, err := reader.ReadString('\n')
					if err != nil {
						fmt.Printf("error reading conn: %s", err)
					}
					//message = strings.TrimSpace(message)
					fmt.Printf("s: %s", message)
				}
			}(conn, i)
			go func(conn net.Conn, strEcho string, i int) {
				//fmt.Printf("Sending message: %s\n", strEcho)
				if _, err = conn.Write(append([]byte(fmt.Sprintf("%s [%d]", strEcho, i)), "\r\n"...)); err != nil {
					println("Write to server failed:", err.Error())
				}
			}(conn, strEcho, i)
		}
	} //
	fmt.Println("end?")
}
