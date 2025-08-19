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

func main(){
	ln,err:= net.Listen("tcp",":8989")
	if err != nil{
		fmt.Println("Server Error:",err)
		return
	}
	for {
		con,err:= ln.Accept()
		if err != nil{
			fmt.Println("Accept Connection Error:",err)
			return
		}
		go functions.HandleConnection(con)
	}
}
