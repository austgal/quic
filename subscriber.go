package main

import (
	"context"
	"log"

	"github.com/quic-go/quic-go"
)

func (c *Connections) handleSubscriber(connection quic.Connection) {
	log.Printf("New subscriber connected: %v\n", connection.RemoteAddr())
	c.addSubscriber(connection)

	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		go func(stream quic.Stream) {
			defer func() {
				log.Printf("Subscriber stream closed: %v\n", stream.StreamID())
				c.removeSubscriber(connection)
			}()

			buf := make([]byte, 1024)
			for {
				n, err := stream.Read(buf)
				if err != nil {
					log.Println(err)
					return
				}

				log.Printf("Received from subscriber %v: %v\n", stream.StreamID(), buf[:n])
			}
		}(stream)
	}
}
