package main

import (
	"context"
	"log"

	"github.com/quic-go/quic-go"
)

func (c *Connections) handlePublisher(connection quic.Connection) {
	log.Printf("New publisher connected: %v\n", connection.RemoteAddr())
	c.addPublisher(connection)
	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		go c.handlePubStream(stream)
		defer func() {
			log.Printf("Publisher stream closed: %v\n", stream.StreamID())
			c.removePublisher(connection)
		}()
	}
}

func (c *Connections) handlePubStream(stream quic.Stream) {

	buf := make([]byte, 1024)
	for {
		if len(c.subscribers) == 0 {
			_, err := stream.Write([]byte("No subscribers are connected"))
			if err != nil {
				log.Println(err)
				return
			}
			break
		}
		n, err := stream.Read(buf)
		log.Println(string(buf[:n]))
		if err != nil {
			log.Println(err)
			return
		}
		c.broadcastMessage(buf[:n])
	}
}

func (c *Connections) broadcastMessage(message []byte) {
	for subscriber := range c.subscribers {
		stream, err := subscriber.OpenStream()
		if err != nil {
			log.Println(err)
			continue
		}
		_, err = stream.Write(message)
		log.Printf("Broadcasting message from publisher: %v\n", string(message))
		if err != nil {
			log.Println(err)
		}
	}
}