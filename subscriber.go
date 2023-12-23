package main

import (
	"context"
	"log"

	"github.com/quic-go/quic-go"
)

const new_sub_msg = "Subscriber connected"

func (c *Connections) handleSubscriber(connection quic.Connection) {
	log.Printf("New subscriber connected: %v\n", connection.RemoteAddr())

	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		c.addSubscriber(stream)
		defer func() {
			c.removeSubscriber(stream)
			log.Printf("subscriber stream closed: %v\n", stream.StreamID())
			stream.Close()
		}()
		go c.informPublishers([]byte(new_sub_msg))
	}
}

func (c *Connections) informPublishers(message []byte) {
	for publisher := range c.publishers {
		_, err := publisher.Write([]byte(message))
		log.Printf("informing publisher : %v\n", string(message))
		if err != nil {
			log.Println(err)
		}
	}
}
