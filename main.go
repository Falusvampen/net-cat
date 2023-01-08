package main

import (
	"fmt"
	"net-cat/server"
	"os"
)

var Port = "8989"

// run the server on the given port number
func main() {
	// go run . h354s
	if len(os.Args) == 1 {
		server.StartServer(Port)
	}
	if len(os.Args) == 2 {
		false := server.CheckValidPort(os.Args[1])
		if false {
			server.StartServer(os.Args[1])
		}
	} else {
		fmt.Println("Too many arguments")
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
}
