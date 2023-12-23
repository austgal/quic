package main

import (
	"context"
	"log"
	"net"

	"github.com/quic-go/quic-go"
)

const pubPort = 6666
const subPort = 6667

func main() {
	connections := &Connections{
		subscribers: make(map[quic.Stream]struct{}),
		publishers:  make(map[quic.Stream]struct{}),
	}

	go startServer(pubPort, connections.handlePublisher)
	go startServer(subPort, connections.handleSubscriber)

	select {}
}

func startServer(port int, handler func(quic.Connection)) {
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: port})

	if err != nil {
		log.Fatal(err)
		return
	}

	tr := quic.Transport{
		Conn: udpConn,
	}
	listener, err := tr.Listen(generateTLSConfig(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()

	log.Printf("Listening on port %v\n", port)
	acceptConnections(*listener, handler)
}

func acceptConnections(listener quic.Listener, handler func(quic.Connection)) {
	for {
		connection, err := listener.Accept(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		go handler(connection)
	}
}
