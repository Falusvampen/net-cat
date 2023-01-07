package server

import (
	"bufio"
	"net"
	"strings"
)

var users []*user

type user struct {
	name string
	conn net.Conn
}

// for client to add username and welcome message
func handleConnection(conn net.Conn) {
	pinguSender(conn)
	clientMessage(conn, "Welcome to the server!\nPlease enter your name: ")
	user := &user{
		name: getName(conn),
		conn: conn,
	}
	users = append(users, user)
	serverMessage("Welcome " + user.name + "!\n")
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		broadcastMessage(conn, message)
	}
	serverMessage(user.name + " has left the server!\n")
	for i, u := range users {
		if u == user {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	conn.Close()
}

func getName(conn net.Conn) string {
	for {
		name, _ := bufio.NewReader(conn).ReadString('\n')
		name = strings.TrimSpace(name)
		if name != "" {
			return name
		}
		conn.Write([]byte("You need a username bro!\nTry again: "))
	}
}
