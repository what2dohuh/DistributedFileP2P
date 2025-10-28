package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	tcpopts := TCPTransportOpts{
		ListenAddr:   "localhost:8080",
		HandshakeFun: NOPHandshake,
		Decoder:      &DefaultDecoder{},
	}
	transport := NewTCPTransport(tcpopts)

	assert.Equal(t, transport.ListenAddr, "localhost:8080")

	assert.Nil(t, transport.ListenAndAccept())
}
