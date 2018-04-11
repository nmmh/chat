```
main.go

types

init()

main()
	go clientManager.start()		
		for
			select
			case reads:			
			case writes:			
			case dels:
		
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
			
		case msgChannel 
            switch scope
            case ALL
                broadcast()
            case P2P
                whisper
            case sender
                sendmessage()
		
		case deadConnections
            senddisconnect()
            dels<deadConnections
```