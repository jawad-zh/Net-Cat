package functions

import (
	"fmt"
	"net"
	"io"
)

// type client struct{
// Name string
// message string
// }
var client []net.Conn
var historique []string
func HandleConnection(con net.Conn  ){
	var Name string
	client = append(client, con)
	// name := client.Name
	defer con.Close()
		msg := `Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    '.       | '' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     '-'       '--'
[ENTER YOUR NAME]: `
	_, err := con.Write([]byte(msg))
	buff:= make([]byte,1024)
	if err != nil {
		fmt.Println("Writing Error:",err)
		return
	}
	n,err:=con.Read(buff)
	 if err!= nil{
		fmt.Print("Server error:",err)
		return
		}else{
		Name = string(buff[:n])
		Brodcast(Name + "is connected\n" ,con)
	}
	for{
		n,err:=con.Read(buff)
		if err == io.EOF{
			Brodcast(Name+"is disconnected\n",con)
			return
		}else if err!= nil{
			fmt.Print("Error 1:",err)
			return
		}
		
		Brodcast("["+Name+"]"+":"+string(buff[:n]),con)
		historique = append(historique, Name,string(buff[:n]))
		if err == nil{
			// fmt.Print(historique)
		}
	}

// fmt.Println(test)
// fmt.Println("------------------")
}

func Brodcast(message string , sender net.Conn){
	for _,c:= range client{
		if c != sender{
			_,err:=c.Write([]byte(message,))
			if err != nil && err != io.EOF{
				fmt.Println("Error: 2",err)
				return
			}
		}
	}
}