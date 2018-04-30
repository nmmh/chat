package main

import (
	"flag"
)

func main() {
	//get config from json config file.
	configuration := Configuration{}
	err := GetConfigFromJSON("../../assets/config.json", &configuration)
	if err != nil {
		panic(err)
	}

	flag.IntVar(&configuration.Port, "port", 6000, "The port to run the Chat server on")
	flag.Parse()

	server, err := NewChatServer(&configuration)
	if err != nil {
		panic(err.Error())
	}
	server.Start()
}
