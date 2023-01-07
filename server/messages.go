package server

import (
	"net"
)

func clientMessage(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

func broadcastMessage(conn net.Conn, message string) {
	if len(message) != 0 {
		for _, user := range users {
			clientMessage(user.conn, message)
		}
	}
}

func serverMessage(message string) {
	for _, user := range users {
		clientMessage(user.conn, message)
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
	"   HZM            CO00",
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
