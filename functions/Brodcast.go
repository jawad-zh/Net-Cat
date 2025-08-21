package functions

import (
	"fmt"
	"io"
	"net"
	"time"
)

func Brodcast(t time.Time, message string, sender net.Conn, client []net.Conn, n string) {
	formatted := "[" + t.Format("2006 01 02 15:04:05") + "]" + " " + message
	for _, c := range client {
		if c != sender {
			_, err := c.Write([]byte(formatted))
			if err == io.EOF {
				fmt.Print(n + "Disconnected")
			}
			if err != nil {
				fmt.Println("Error: 2", err)
				return
			}
		}
	}
}
