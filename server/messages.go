package server

import (
	"net"
)

// Function that writes to one client
func clientMessage(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

// Sends a message from a client to all clients, locks it, appends it to the history and then unlocks it
func broadcastMessage(conn net.Conn, user string, message string) {
	finalMessage := getTime() + user + message
	if len(message) != 3 {
		// add message to history
		mu.Lock()
		History = append(History, finalMessage)
		mu.Unlock()
		for _, user := range users {
			clientMessage(user.conn, finalMessage)
		}
	}
}

// Sends a message from the server to all clients, locks it, appends it to the history and then unlocks it
func serverMessage(message string) {
	mu.Lock()
	History = append(History, getTime()+message)
	mu.Unlock()
	for _, user := range users {
		clientMessage(user.conn, getTime()+message)
	}
}

// sends the pingu to the client
func pinguSender(conn net.Conn) {
	for _, e := range pinguAlive {
		clientMessage(conn, e+"\n")
	}
}
