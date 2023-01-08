package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Shared variables
var (
	users   []*User
	History = []string{}
	mu      sync.Mutex
)

type User struct {
	name string
	conn net.Conn
}

// for client to add username and welcome message
func handleConnection(conn net.Conn) {
	// First send welcome message and ask for username to client
	pinguSender(conn)
	clientMessage(conn, "Welcome to the server!\nWrite /help to see the available commands\nPlease enter your name: ")
	user := &User{
		name: getName(conn),
		conn: conn,
	}
	//add user to the list of users and send message to the other users
	mu.Lock()
	users = append(users, user)
	mu.Unlock()
	// send the history to the client
	for _, message := range History {
		clientMessage(conn, message)
	}
	serverMessage(user.name + " has joined the server!\n")
	fmt.Println(getTime() + user.name + " has joined the server" + "!")
	// Continuously read messages from the client and broadcast them
InputLoop:
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		switch {
		case message == "/name\n":
			changeName(conn, user, message)
			continue
		case message == "/exit\n":
			break InputLoop
		case message == "/help\n":
			clientMessage(conn, "Commands:\n/name - change your name\n/exit - exit the server\n/help - show this message\n")
			continue
		default:
			broadcastMessage(conn, user.name, ": "+message)
		}
	}
	// remove user from the list of users and close the connection
	removeClient(conn, user)
}

// Function to remove a client from the list of users and close the connection to the client and broadcast it to the other users
func removeClient(conn net.Conn, user *User) {
	serverMessage(user.name + " has left the server!\n")
	fmt.Println(getTime() + user.name + " has left the server!")
	mu.Lock()
	for i, u := range users {
		if u.conn == conn {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	fmt.Println()
	mu.Unlock()
	conn.Close()
}

// Change the name of the user and broadcast it to the other users
func changeName(conn net.Conn, user *User, name string) {
	mu.Lock()
	clientMessage(conn, "Please enter your new name: ")
	oldName := user.name
	user.name = getName(conn)
	mu.Unlock()
	broadcastMessage(conn, oldName, " has changed their name to "+user.name+"!\n")
}

// Get the name of the user and make sure it isn't empty
func getName(conn net.Conn) string {
	for {
		name, _ := bufio.NewReader(conn).ReadString('\n')
		name = strings.TrimSpace(name)
		if name != "" {
			return randomUserNameColor(name)
		}
		conn.Write([]byte("You need a username bro!\nTry again: "))
	}
}

func randomUserNameColor(username string) string {
	randSeed := time.Now().UnixNano()
	randSeed = randSeed % int64(230)
	return "\033[38;5;" + Itoa(int(randSeed)) + "m" + username + "\033[0m"
}
