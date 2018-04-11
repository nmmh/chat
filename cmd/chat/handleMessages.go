package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// this really belongs in ClientManager.go - i had to make the func below a method of clientManger so it could "see" the clientMananger channels
// ...I had intended for this function to be a pure networks and comms funciton. Because of channel scope this concept became a bit muddled in the end.
//... eg. client manager performs the broadcast because I needed to range through the clients.

//dont want to pass conn here but will for the moment.
func (cm *ClientManager) handleMessages(conn net.Conn, cs *ClientState, reader *bufio.Reader) {
	for {
		incoming, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		incoming = strings.NewReplacer("\r", "", "\n", "").Replace(incoming)

		if strings.HasPrefix(incoming, "/") {
			if strings.HasPrefix(incoming, "/bye") {
				break
			}
			if strings.HasPrefix(incoming, "/list") {
				usernames := cm.ReadAll()
				ul, _ := cm.FormatUserList(usernames)
				msgChannel <- &message{cs.username, "CHANOP", "SENDERONLY", fmt.Sprintf("%s\r\n", ul)}
				continue
			}

			if strings.HasPrefix(incoming, "/help") {
				msgChannel <- &message{cs.username, "CHANOP", "SENDERONLY", fmt.Sprintf("%s\r\n", getWelcome())}
				continue
			}
			if strings.HasPrefix(incoming, "/username") {
				prevUsername := cs.username
				newUsername := strings.NewReplacer("/username ", "", "\r", "", "\n", "", "anonymous", "").Replace(incoming)
				if len(newUsername) > 0 && strings.Index(newUsername, "/") == -1 {
					usernames := cm.ReadAll()
					if uiu, _ := cm.UsernameInUse(usernames, newUsername); !uiu {
						if cm.Write(conn, newUsername) {
							cs.username = newUsername
							msgChannel <- &message{newUsername, "CHANOP", "ALL", fmt.Sprintf(" * %s changed username to [%s]\r\n", prevUsername, newUsername)}
						}
						continue
					} else {
						msgChannel <- &message{prevUsername, "CHANOP", "SENDERONLY", fmt.Sprintf(" * %s failed to change username to \"%s\"\r\n", prevUsername, newUsername)}
						continue
					}
				} else {
					msgChannel <- &message{prevUsername, "CHANOP", "SENDERONLY", fmt.Sprintf(" * %s failed to change username to \"%s\"\r\n", prevUsername, newUsername)}
					continue
				}
			}
			//all the valid commands have been processed and this one was not understood.
			msgChannel <- &message{cs.username, "CHANOP", "SENDERONLY", fmt.Sprintf(" * %s issued unrecognised command \"%s\"\r\n", cs.username, incoming)}
			continue
		} else if strings.HasPrefix(string(incoming), "@") {
			//extract username and message
			if strings.Index(incoming, " ") > 1 && len(incoming) > 3 {
				usernames := cm.ReadAll()
				whisperToUser := incoming[1:strings.Index(incoming, " ")]
				whisperMsg := incoming[strings.Index(incoming, " ")+1 : len(incoming)]
				if uiu, _ := cm.UsernameInUse(usernames, whisperToUser); uiu {
					msgChannel <- &message{whisperToUser, "WHISPER", "SENDERONLY", fmt.Sprintf("[%s] whispers> %s\r\n", cs.username, whisperMsg)}
					continue
				} else {
					msgChannel <- &message{cs.username, "WHISPER", "SENDERONLY", fmt.Sprintf(" * \"%s\" is unknown user\r\n", whisperToUser)}
					continue
				}
			} else {
				msgChannel <- &message{cs.username, "WHISPER", "SENDERONLY", fmt.Sprintf(" * %s issued unrecognised command \"%s\"\r\n", cs.username, incoming)}
				continue
			}
		}

		msgChannel <- &message{cs.username, "NORMAL", "ALLEXCEPTSENDER", fmt.Sprintf("[%s]> %s\r\n", cs.username, incoming)}
	}
	// When we encouter `err` reading, send this
	// connection to `deadConnections` for removal.
	//

	deadConnections <- conn
}
