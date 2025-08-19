package functions

import (
	"fmt"
	"net"
	"io"
	"time"
)

// type client struct{
// Name string
// message string
// }

var historique []string
var users []string
func HandleConnection(con net.Conn  ,client []net.Conn){
	var Name string
	
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
			for _,r:= range users{
				if Name==r{
					con.Write([]byte("the name is token try other one"))
					return
				}else{
					users = append(users, Name)
				}
			}		
		Brodcast(time.Now(),Name + "is connected\n" ,con ,client)
		
	}
	for{
		n,err:=con.Read(buff)
		if err == io.EOF{
			Brodcast(time.Now(),Name+"is disconnected\n",con ,client)
			return
		}else if err!= nil{
			fmt.Print("Error 1:",err)
			return
		}
		
		Brodcast(time.Now(),"["+Name+"]"+":"+string(buff[:n]),con ,client)
		historique = append(historique, Name,string(buff[:n]))
		if err == nil{
			// fmt.Print(historique)
		}
	}

// fmt.Println(test)
// fmt.Println("------------------")
}

func Brodcast(t time.Time ,message string , sender net.Conn,client []net.Conn){
	formatted := fmt.Sprintf("[%s] %s", t.Format("15:04:05"), message)
	for _,c:= range client{
		if c != sender{
			_,err:=c.Write([]byte(formatted))
			if err != nil && err != io.EOF{
				fmt.Println("Error: 2",err)
				return
			}
		}
	}
}