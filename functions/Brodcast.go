package functions

import (
	"fmt"
	"net"
	"time"
)

func Brodcast(t time.Time, message string, sender net.Conn, clients *[]net.Conn, n string) {
	formatted := "[" + t.Format("2006 01 02 15:04:05") + "] " + message

	for i := 0; i < len(*clients); i++ {
		c := (*clients)[i]
		if c != sender {
			_, err := c.Write([]byte(formatted))
			if err != nil {
				fmt.Println(n + " Disconnected")
				*clients = append((*clients)[:i], (*clients)[i+1:]...)
				i-- 
				continue
			}
		}
	}
}
