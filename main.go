package main

import (
	"fmt"
	"functions/functions"
	"net"
)

// type client struct{
// Name string
// message string
// }
var client []net.Conn
func main(){

	ln,err:= net.Listen("tcp",":8989")
	if err != nil{
		fmt.Println("Server Error:",err)
		return
	}
	for len(client)<=10{
		con,err:= ln.Accept()
		if err != nil{
			fmt.Println("Accept Connection Error:",err)
			return
		}
		client = append(client, con)
		go functions.HandleConnection(con,client)
	}
}
