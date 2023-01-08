package server

import (
	"fmt"
	"log"
	"net"
)

// start the server and connect to the server
func StartServer(Port string) error {
	fmt.Println("Server started port:" + Port)
	listener, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// listen on the server and connect to the server
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if len(users) < 10 {
			go handleConnection(conn)
		} else {
			clientMessage(conn, "Server is full, try again later!\n")
			conn.Close()
		}
	}
}
