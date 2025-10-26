package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn
	// if we dial outbound = true
	// if we accept inbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr   string
	HandshakeFun HandshakeFunction
	Decoder      Decorder
}
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		peers:            make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFun(peer); err != nil {
		fmt.Println("Handshake failed:", err)
		conn.Close()
		return
	}

	t.mu.Lock()
	t.peers[conn.RemoteAddr()] = peer
	t.mu.Unlock()

	msg := &Message{}

	fmt.Printf("New connection from: %+v\n", peer)
	fmt.Printf("Connections : %+v\n", len(t.peers))

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Println("Decode error:", err)
			continue
		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("Received message: %+v\n", msg)
		fmt.Printf("Received message: %+v\n", string(msg.Payload))
	}

}
