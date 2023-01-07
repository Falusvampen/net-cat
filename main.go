package main

import (
	"net-cat/server"
	"os"
)

var Port = "8989"
// run the server on the given port number
func main() {
	// go run . 8546
	if len(os.Args) > 1 {
		Port = os.Args[1]
	}
	server.StartServer(Port)
}
