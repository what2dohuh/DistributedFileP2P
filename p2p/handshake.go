package p2p

type HandshakeFunction func(any) error 

func NOPHandshake(any) error{return nil}