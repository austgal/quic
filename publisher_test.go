package main

import (
	"errors"
	"testing"

	"github.com/quic-go/quic-go"
	"github.com/stretchr/testify/mock"
)

func TestBroadcastMessage(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Connection]struct{}),
		publishers:  make(map[quic.Connection]struct{}),
	}

	mockSubscriber1 := new(MockConnection)
	mockSubscriber2 := new(MockConnection)

	connections.subscribers[mockSubscriber1] = struct{}{}
	connections.subscribers[mockSubscriber2] = struct{}{}

	mockStream1 := new(MockStream)
	mockStream2 := new(MockStream)

	mockSubscriber1.On("OpenStream").Return(mockStream1, nil).Once()
	mockSubscriber2.On("OpenStream").Return(mockStream2, nil).Once()

	mockStream1.On("Write", []byte("Notifying subscribers!")).Return(22, nil).Once()
	mockStream2.On("Write", mock.Anything).Return(0, errors.New("write error")).Once()

	connections.broadcastMessage([]byte("Notifying subscribers!"))

	mockSubscriber1.AssertExpectations(t)
	mockSubscriber2.AssertExpectations(t)
	mockStream1.AssertExpectations(t)
	mockStream2.AssertExpectations(t)
}
