package server

import (
	"fmt"
	"time"
)

func getTime() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
}

// Itoa function because we aren't allowed to use strconv
func Itoa(i int) string {
	if i == 0 {
		return "0"
	}
	res := ""
	for i > 0 {
		res = string((i%10 + 48)) + res
		i /= 10
	}
	if i < 0 {
		return "-" + Itoa(-i)
	} else {
		return res
	}
}

func validAtoi(s string) (int, bool) {
	res := 0
	if s[0] == '-' {
		fmt.Println("Negative numbers are not allowed!")
		return 0, false
	}
	for _, e := range s {
		if e >= '0' && e <= '9' {
			res = res*10 + int(e-48)
		} else {
			fmt.Println("Only numbers are allowed!")
			return 0, false
		}
	}
	return res, true
}

func CheckValidPort(port string) bool {
	nbr, err := validAtoi(port)
	if !err {
		return false
	}
	if nbr < 0 || nbr > 65535 {
		fmt.Println("Port number must be between 0 and 65535")
		return false
	}
	return true
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
