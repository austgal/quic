package mocks

import (
	"context"
	"net"

	"github.com/quic-go/quic-go"
	"github.com/stretchr/testify/mock"
)

// MockConnection is a mock implementation for the Connection interface
type MockConnection struct {
	mock.Mock
}

func (m *MockConnection) AcceptStream(ctx context.Context) (quic.Stream, error) {
	args := m.Called(ctx)
	return args.Get(0).(quic.Stream), args.Error(1)
}

func (m *MockConnection) AcceptUniStream(ctx context.Context) (quic.ReceiveStream, error) {
	args := m.Called(ctx)
	return args.Get(0).(quic.ReceiveStream), args.Error(1)
}

func (m *MockConnection) OpenStream() (quic.Stream, error) {
	args := m.Called()
	return args.Get(0).(quic.Stream), args.Error(1)
}

func (m *MockConnection) OpenStreamSync(ctx context.Context) (quic.Stream, error) {
	args := m.Called(ctx)
	return args.Get(0).(quic.Stream), args.Error(1)
}

func (m *MockConnection) OpenUniStream() (quic.SendStream, error) {
	args := m.Called()
	return args.Get(0).(quic.SendStream), args.Error(1)
}

func (m *MockConnection) OpenUniStreamSync(ctx context.Context) (quic.SendStream, error) {
	args := m.Called(ctx)
	return args.Get(0).(quic.SendStream), args.Error(1)
}

func (m *MockConnection) LocalAddr() net.Addr {
	args := m.Called()
	return args.Get(0).(net.Addr)
}

func (m *MockConnection) RemoteAddr() net.Addr {
	args := m.Called()
	return args.Get(0).(net.Addr)
}

func (m *MockConnection) CloseWithError(code quic.ApplicationErrorCode, reason string) error {
	args := m.Called(code, reason)
	return args.Error(0)
}

func (m *MockConnection) ConnectionState() quic.ConnectionState {
	args := m.Called()
	return args.Get(0).(quic.ConnectionState)
}

func (m *MockConnection) Context() context.Context {
	args := m.Called()
	return args.Get(0).(context.Context)
}

func (m *MockConnection) SendMessage(message []byte) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockConnection) ReceiveMessage() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockConnection) ReceiveDatagram(context.Context) ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockConnection) SendDatagram(_ []byte) error {
	args := m.Called()
	return args.Error(0)
}

// type mockConnection struct{}
// type mockConnectionState struct {
// 	// Define fields as needed for your mock implementation
// }

// func (m *mockConnection) AcceptStream(context.Context) (quic.Stream, error) { return nil, nil }
// func (m *mockConnection) AcceptUniStream(context.Context) (quic.ReceiveStream, error) {
// 	return nil, nil
// }
// func (m *mockConnection) OpenStream() (quic.Stream, error) {
// 	return nil, nil
// }
// func (m *mockConnection) OpenStreamSync() (quic.Stream, error) {
// 	return nil, nil
// }
// func (m *mockConnection) OpenUniStream() (quic.Stream, error) {
// 	return nil, nil
// }
// func (m *mockConnection) OpenUniStreamSync(context.Context) (quic.Stream, error) {
// 	return nil, nil
// }
// func (m *mockConnection) LocalAddr() net.Addr  { return nil }
// func (m *mockConnection) RemoteAddr() net.Addr { return nil }
// func (m *mockConnection) CloseWithError(quic.ApplicationErrorCode, string) error {
// 	return nil
// }

// func (m *mockConnection) ConnectionState() quic.ConnectionState {
// 	// Return an instance of your mockConnectionState type
// 	return &mockConnectionState{}
// }

// var _ quic.ConnectionState = &mockConnectionState{}

// // Implement the quic.ConnectionState interface for mockConnectionState
// func (m *mockConnectionState) ConnectionState() quic.ConnectionState {
// 	return m
// }
