package server

import (
	"fmt"
	"log"
	"net"
)

// start the server and connect to the server
func StartServer(Port string) error {
	fmt.Println("Server started port:" + Port)
	listener, err := net.Listen("tcp", "localhost:"+Port)
	if err != nil {
		log.Fatal(err)
	}
	// listen on the server and connect to the server
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if len(users) < 10 {
			fmt.Println("New connection from " + conn.RemoteAddr().String())
			go handleConnection(conn)
			// print the connected clients to the server
			fmt.Println("Connected clients: ", len(users)+1)
		} else {
			clientMessage(conn, "Server is full, try again later!\n")
			conn.Close()
		}
	}
}
