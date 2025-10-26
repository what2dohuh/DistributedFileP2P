package p2p

type Peer interface {

}

//Handles communication between nodes
type Transport interface {
	ListenAndAccept() error
}