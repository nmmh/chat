
# Chat - a golang concurrent tcp chat server
## This is a basic chat server that uses goroutines. A telnet client can be used to connect.
### This program is an __attempt__ to demonstrate the maintenance of state between goroutines using channels. This is a learning excercise.

>"Do not communicate by sharing memory; instead, share memory by communicating."

Pfff... very tricky.

So I wanted to protect a map from problems arising from being accessed by several different go routines by using channels. 
It could probably be better written with a RWMUTEX.   
I considered using the sync map included in GO 1.9 but was terrified of the warning regarding ranging through the map - no guarantee of a consistent map - I dont really know specifically what that means, but I knew I needed to range through the map and so stayed away...

**I am new to GOLANG.**  
**Postive/constructive criticism will be gratefully received**

TODO:  
* Close channels when the app stops. (I assume the channels in here are a concurrent travesty in terms of leaks etc.)
* ~~stop over using channels~~
* Clean up messageHandler ~~lots of boilerplate messages~~ still lots of conditions
* ~~Resolve ownership of the channels (Its not clear whether a separate set of network channels should be used)~~
* ~~JSON Config file~~
* ~~pass server port with override~~
* tests
* testing on other OSs - currently windows (I am sure there are CRLF issues)
* other things

### Install:  
____________
```
$ go get github.com/nmmh/chat
$ cd %GO_PATH%/src/github.com/nmmh/chat/cmd/chat
$ go build; go install
```
____________
### Server Usage:  
____________
```
$ ./chat
```
____________

### Telnet Client Usage:  
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

### General structure.  (Should be more staight forward now from ChatServer.Start() )
```
go s.sendMessages() //has the channels	
for { 
	conn, e := s.lsnr.Accept() 
	if e != nil { 
		continue 
	} 
	go s.handleClient(conn) 
}
```
___________			
### Screenshot
![alt text](https://github.com/nmmh/chat/blob/master/assets/chat_screenshot.JPG?raw=true "screenshot of server and 3 client sessions")