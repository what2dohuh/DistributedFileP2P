package p2p
import(
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestTcpTransport(t *testing.T) {
	transport := NewTCPTransport("localhost:8080")

	assert.Equal(t, transport.listenAddr,"localhost:8080" )

	assert.Nil(t,transport.ListenAndAccept())
	select {}
}