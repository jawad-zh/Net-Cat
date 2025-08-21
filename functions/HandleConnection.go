package functions

import (
	"fmt"
	"io"
	"net"
	"time"
)

// type client struct{
// Name string
// message string
// }

// var historique []string

func HandleConnection(con net.Conn) {
	var Name string
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

	buff := make([]byte, 1024)
	flag:= true
	for flag {

		_, err := con.Write([]byte(msg))
	
	if err != nil {
		fmt.Println("Writing Error:", err)
		return
	}
	
	n, err := con.Read(buff)
	if err != nil {
		fmt.Print("Server error:", err)
		return
	} 
		Name = string(buff[:n])
		if !ChekConn(Name){
			con.Write([]byte("The name is Token Try an other one\n"))
			continue
		}else{
			flag= false
			Clients = append(Clients, Name)
		}
		}

		Brodcast(time.Now(), Name+"is connected\n", con, AddClients)

	
	for {
		n, err := con.Read(buff)
		if err == io.EOF {
			Brodcast(time.Now(), Name+"is disconnected\n", con, AddClients)
			return
		} else if err != nil {
			fmt.Print("Error 1:", err)
			return
		}

		Brodcast(time.Now(), "["+Name+"]"+":"+string(buff[:n]), con, AddClients)
		// historique = append(historique, Name, string(buff[:n]))
	}

	// fmt.Println(test)
	// fmt.Println("------------------")
}

func Brodcast(t time.Time, message string, sender net.Conn, client []net.Conn) {
	formatted := t.Format("15:04:05") + message
	for _, c := range client {
		if c != sender {
			_, err := c.Write([]byte(formatted))
				if err != nil {
				fmt.Println("Error: 2", err)
				return
			}
		}
	}
}
