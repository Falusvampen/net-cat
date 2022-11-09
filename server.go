package main

import (
	"log"
)

const (
	CONN_TYPE = "tcp"

	MAX_CLIENTS = 10

	CMD_PREFIX = "/"
	CMD_CREATE = CMD_PREFIX + "create"
	CMD_JOIN   = CMD_PREFIX + "join"
	CMD_LEAVE  = CMD_PREFIX + "leave"
	CMD_LIST   = CMD_PREFIX + "list"
	CMD_QUIT   = CMD_PREFIX + "quit"
	CMD_HELP   = CMD_PREFIX + "help"
	CMD_NAME   = CMD_PREFIX + "name"

	CLIENT_NAME = "User"
	SERVER_NAME = "Server"

	MSG_WELCOME = "Welcome to the chat server!"

	ERROR_PREFIX = "Error: "
	ERROR_SEND   = ERROR_PREFIX + "No messages allowed in the lobby. Join a chat room"
	ERROR_JOIN   = ERROR_PREFIX + "You are already in a chat room"
	ERROR_LEAVE  = ERROR_PREFIX + "You are not in a chat room"
	ERROR_CREATE = ERROR_PREFIX + "You are already in a chat room or the room already exists"

	NOTICE_PREFIX      = "Notice: "
	NOTICE_ROOM_JOIN   = NOTICE_PREFIX + "You have joined the chat room"
	NOTICE_ROOM_LEAVE  = NOTICE_PREFIX + "You have left the chat room"
	NOTICE_ROOM_CREATE = NOTICE_PREFIX + "You have created the chat room"
	NOTICE_ROOM_LIST   = NOTICE_PREFIX + "Chat rooms: "
	NOTICE_DELETE      = NOTICE_PREFIX + "the chat room has been deleted"

	MSG_CONNECT = "Welcome to the server! Type \"/help\" to get a list of commands.\n"
	MSG_FULL    = "The server is full. Try again at a later time.\n"
)

// The lobby keeps track of chat rooms and clients
type Lobby struct {
	clients  []*Client
	rooms    map[string]*Room
	incoming chan *Message
	join     chan *Client
	leave    chan *Client
	delete   chan *Room
}

func NewLobby() *Lobby {
	lobby := &Lobby{
		clients:  make([]*Client, 0),     //creates a slice of clients
		rooms:    make(map[string]*Room), //creates a map of rooms
		incoming: make(chan *Message),    //creates a channel for incoming messages
		join:     make(chan *Client),     //creates a channel for clients joining rooms
		leave:    make(chan *Client),     //creates a channel for clients leaving rooms
		delete:   make(chan *Room),       //creates a channel for rooms being deleted
	}
	lobby.Listen() //starts listening to the lobby's channels to detect activity on the server
	return lobby
}

// listens to the lobby's channels and handles the messages
func (lobby *Lobby) Listen() {
	go func() {
		for {
			select {
			case message := <-lobby.incoming: //listen for incoming messages and parses them
				lobby.Parse(message)
			case client := <-lobby.join: //listens to see if clients want to join a room
				lobby.Join(client)
			case client := <-lobby.leave: //listens to see if clients want to leave a room
				lobby.Leave(client)
			case room := <-lobby.delete: //listens to see if rooms are being deleted
				lobby.Delete(room)
			}
		}
	}()
}

func (lobby *Lobby) Join(client *Client) {
	if len(lobby.clients) >= MAX_CLIENTS { //checks to see if the server is full
		client.Quit()
		return
	}
	lobby.clients = append(lobby.clients, client) //adding the newest client to the list of clients
	client.outgoing <- MSG_WELCOME                //sending the welcome message to the client
	go func() {
		for messages := range client.incoming {
			lobby.incoming <- messages
		}
		lobby.leave <- client
	}()
}

func (lobby *Lobby) Leave(client *Client) {
	if client.room != nil { //removes the client from the chatrom if the room is not empty
		client.room.Leave(client)
	}
	for i, c := range lobby.clients { //looks for the client to remove
		if client == c {
			lobby.clients = append(lobby.clients[:i], lobby.clients[i+1]...)
		}
	}
	close(client.outgoing)
	log.Println("Closed the outgoing channel of the client")
}

// deletes the room
func (lobby *Lobby) DeleteChatRoom(room *Room) {
	room.Delete()
	delete(lobby.rooms, room.name)
	log.Println("succsesfully deletet the chat room")
}
