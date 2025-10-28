# TCPP2P

## Overview

TCPP2P is a lightweight, foundational Go library for building peer-to-peer (P2P) applications over TCP. It provides a simple and extensible framework for nodes to establish connections, perform handshakes, and exchange messages. The design emphasizes modularity through interfaces, allowing developers to easily customize handshake logic and message encoding.

## Features

*   **TCP Transport Layer:** Provides a transport layer for P2P communication using standard TCP sockets.
*   **Modular Design:** Utilizes interfaces for `Transport`, `Peer`, and `Decoder` to allow for easy extension and customization.
*   **Pluggable Handshake:** Define custom handshake logic by implementing the `HandshakeFunction`. A no-op handshake is provided by default.
*   **Customizable Encoding:** Swap out decoding logic by implementing the `Decoder` interface. `DefaultDecoder` (raw bytes) and `GOBDecoder` are included.
*   **Concurrent Connections:** Handles multiple inbound peer connections concurrently using goroutines.

## Getting Started

### Prerequisites

*   Go (version 1.24 or later)

### Installation

Clone the repository to your local machine:
```sh
git clone https://github.com/what2dohuh/tcpp2p.git
cd tcpp2p
```

### Running the Server

This project uses a `Makefile` to simplify common tasks.

To build the application:
```sh
make build
```

To build and run the server, which will start listening for connections on port `:8080`:
```sh
make run
```

To run the test suite:
```sh
make test
```

## Usage

You can integrate TCPP2P into your own project as a module. The following example shows how to start a TCP transport listener.

```go
package main

import (
	"fmt"
	"github.com/what2dohuh/distributedFile/p2p"
)

func main() {
	// Configure the TCP transport options
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:   ":8080",
		HandshakeFun: p2p.NOPHandshake,
		Decoder:      &p2p.DefaultDecoder{},
	}

	// Create a new TCP transport
	tr := p2p.NewTCPTransport(tcpOpts)

	// Start listening for and accepting incoming connections
	err := tr.ListenAndAccept()
	if err != nil {
		fmt.Println("Error starting TCP Transport:", err)
		return
	}

	// Block the main thread to keep the server running
	select {}
}
```
When a peer connects, the server will handle the connection, perform the handshake, and decode incoming messages, printing them to the console.

## Project Structure

The core logic is located in the `p2p/` directory:

*   **`transport.go`**: Defines the core `Transport` and `Peer` interfaces for network communication.
*   **`tcp_transport.go`**: Implements the `Transport` interface for TCP. It manages peer connections and message loops.
*   **`message.go`**: Defines the `RPC` struct, which represents a message received from a peer, containing the payload and sender's address.
*   **`encoding.go`**: Defines the `Decoder` interface for message serialization and provides `DefaultDecoder` (for raw bytes) and `GOBDecoder` implementations.
*   **`handshake.go`**: Defines the `HandshakeFunction` type, allowing for custom logic to be executed when a new peer connects.
