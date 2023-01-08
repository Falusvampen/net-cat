package server

import (
	"net"
	"time"
)

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

// Sends a message from the server to the client, locks it, appends it to the history and then unlocks it
func serverMessage(message string) {
	mu.Lock()
	History = append(History, getTime()+message)
	mu.Unlock()
	for _, user := range users {
		clientMessage(user.conn, getTime()+message)
	}
}

var pinguAlive = []string{
	"         _nnnn_",
	"        dAASSMMb",
	"       @p~qp~~qMb",
	"       M|@||@) M|",
	"       @,----.JM|",
	"      JS^\\__/  qKL",
	"     dZP        ciaw",
	"    dZP          c42g",
	"   BIG    helo    NICE",
	"   HZM            BO00",
	"   LoL            o0CC",
	" __| \".        |\\dS\"qML",
	" |    `.       | `' \\Zq",
	"_)      \\.___.,|     .'",
	"\\____   )RUUDNA|   .'",
	"     `-'       `--'",
}

func pinguSender(conn net.Conn) {
	for _, e := range pinguAlive {
		clientMessage(conn, e+"\n")
	}
}

func getTime() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
}
