package main

import (
	"fmt"

	"github.com/what2dohuh/distributedFile/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:   ":8080",
		HandshakeFun: p2p.NOPHandshake,
		Decoder:      &p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	err := tr.ListenAndAccept()
	if err != nil {
		fmt.Println("Error starting TCP Transport:", err)
		return
	}
	select {}
}
