package functions

import (
	"fmt"
	"net"
)

type client struct{
Name string
message string
}

func HandleConnection(con net.Conn  ){
	
	// name := client.Name
	defer con.Close()
	_,err:=con.Write([]byte("your name:"))
	buff:= make([]byte,1024)
	if err != nil {
		fmt.Println("Writing Error:",err)
		return
	}
	n,err:=con.Read(buff)
	if err != nil{
		fmt.Println("the client is disconected")
		return
	}else{
	}


}