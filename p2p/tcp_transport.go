package p2p

import (
	"net"
	"sync"
	"fmt"
)



type TCPTransport struct {
	listenAddr 	string
	listener   	net.Listener
	mu 	 	   	sync.RWMutex
	peers    	map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listenAddr,
		peers:      make(map[net.Addr]Peer),
	}	
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop(){

	conn, err := t.listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	go t.handleConn(conn)
}

func (t *TCPTransport) handleConn(conn net.Conn)  {
	fmt.Println("New connection from:", conn)
}