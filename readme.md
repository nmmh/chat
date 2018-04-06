
# Chat - a golang chat tcp server
## This is a basic chat server that uses goroutines. A telnet client can be used to connect.
### This program is an __attempt__ to demonstrate the maintenance of state between goroutines using channels. This is a learning excercise.

>"Do not communicate by sharing memory; instead, share memory by communicating."

Pfff... very tricky from what I am more used to.

Anyway, this monster could probably be better written with a RWMUTEX.   
I considered using the sync map included in GO 1.9 but was terrified of the warning regarding ranging through the map - no guarantee of a consistent map - I dont really know specifically what that means, but I knew I needed to range through the map and so stayed away...

**I am new to GOLANG.**  
**Postive/constructive criticism will be gratefully received**

TODO:  
* Close channels when the app stops. (I assume the channels in here are a concurrent travesty)
* Clean up messageHandler
* Resolve ownership of the channels (Its not clear whether a separate set of network channels should be used)
* other things

### Install:  
____________
```
$ go get github.com/nmmh/chat
$ go build
```
____________
### Server Usage:  
____________
```
$ ./chat
```
____________

### Client Usage:  
____________
```
$ telnet 127.0.0.1 6000
```
____________
```
* Usage:
*     just type stuff and press enter to chat.
*
*  Cmds:
*    /username blah (change username to blah)
*    @blah hello (whispers "hello" to blah)
*    /list (show user list)
*    /bye (quits!)
*    /help (displays this msg)
```
____________

### General structure.  
```
main.go

types

init()

main()
	go clientManager.start()		
		select
		case reads:
		case readallvals:
		case writes:
		case msgsforbroadcast:
		case kills:
	
	start tcp server
	
	go func()
		for 
			accept.tcp conn			
			submit on newConnections <		
		
	for 
		select
		case newConnections
			get clientsate for this connection.
			
			go messageHandler()
			
		case msgsForBroacast < msgChannel 
		
		case kills < deadConnections
```
___________			
### Screenshot
![alt text](https://github.com/nmmh/chat2/blob/master/chat_screenshot.JPG?raw=true "screenshot of server and 3 client sessions")