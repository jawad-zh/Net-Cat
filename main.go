package main

import (
	"fmt"
	"functions/functions"
	"net"
)

type client struct{
Name string
message string
}

func main(){
	ln,err:= net.Listen("tcp",":1234")
	if err != nil{
		fmt.Println("Server Error:",err)
		return
	}
	for{
		conn,err:= ln.Accept()
		if err != nil{
			fmt.Println("Error:",err)
			return
		}
	
go functions.HandleConnection(conn)
	}
}
