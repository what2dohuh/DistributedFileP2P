package main
import (	
		"fmt"
		"github.com/what2dohuh/distributedFile/p2p"
		)

func main() {
	tr := p2p.NewTCPTransport(":8080")
	err := tr.ListenAndAccept()
	if err != nil {
		fmt.Println("Error starting TCP Transport:", err)
		return
	}
	select {}
}