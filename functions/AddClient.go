package functions

import (
	"net"
	"sync"
)
var mu sync.Mutex
var AddClients []net.Conn
func AddClient(con net.Conn){
	mu.Lock()
	AddClients = append(AddClients, con)
	mu.Unlock()
}