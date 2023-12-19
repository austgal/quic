package main

import (
	"sync"

	"github.com/quic-go/quic-go"
)

type Connections struct {
	subscribers map[quic.Connection]struct{}
	publishers  map[quic.Connection]struct{}
	mu          sync.Mutex
}

func (c *Connections) addConnection(connection quic.Connection, connectionMap map[quic.Connection]struct{}) {
	c.mu.Lock()
	connectionMap[connection] = struct{}{}
	c.mu.Unlock()
}

func (c *Connections) removeConnection(connection quic.Connection, connectionMap map[quic.Connection]struct{}) {
	c.mu.Lock()
	delete(connectionMap, connection)
	c.mu.Unlock()
}

func (c *Connections) addSubscriber(connection quic.Connection) {
	c.addConnection(connection, c.subscribers)
}

func (c *Connections) addPublisher(connection quic.Connection) {
	c.addConnection(connection, c.publishers)
}

func (c *Connections) removeSubscriber(connection quic.Connection) {
	c.removeConnection(connection, c.subscribers)
}

func (c *Connections) removePublisher(connection quic.Connection) {
	c.removeConnection(connection, c.publishers)
}
