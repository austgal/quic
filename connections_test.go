package main

import (
	"testing"

	"github.com/quic-go/quic-go"
	"github.com/stretchr/testify/assert"
)

func TestAddSubscriber(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Connection]struct{}),
		publishers:  make(map[quic.Connection]struct{}),
	}

	connection := &MockConnection{}
	connections.addSubscriber(connection)

	assert.Contains(t, connections.subscribers, connection)
}

func TestAddPublisher(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Connection]struct{}),
		publishers:  make(map[quic.Connection]struct{}),
	}

	connection := &MockConnection{}
	connections.addPublisher(connection)

	assert.Contains(t, connections.publishers, connection)
}

func TestRemoveSubscriber(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Connection]struct{}),
		publishers:  make(map[quic.Connection]struct{}),
	}

	connection := &MockConnection{}
	connections.subscribers[connection] = struct{}{}
	connections.removeSubscriber(connection)

	assert.NotContains(t, connections.subscribers, connection)
}

func TestRemovePublisher(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Connection]struct{}),
		publishers:  make(map[quic.Connection]struct{}),
	}

	connection := &MockConnection{}
	connections.publishers[connection] = struct{}{}
	connections.removePublisher(connection)

	assert.NotContains(t, connections.publishers, connection)
}
