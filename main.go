package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"net"

	"github.com/quic-go/quic-go"
)

const pubPort = 6666
const subPort = 6667

// TODO: move to separate file
// TODO: maybe add generic func

func main() {
	connections := &Connections{
		subscribers: make(map[quic.Connection]struct{}),
		publishers:  make(map[quic.Connection]struct{}),
	}

	go startServer(pubPort, connections.handlePublisher)
	go startServer(subPort, connections.handleSubscriber)

	select {}
}

func startServer(port int, handler func(quic.Connection)) {
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: port})

	if err != nil {
		log.Fatal(err)
	}

	tr := quic.Transport{
		Conn: udpConn,
	}
	listener, err := tr.Listen(generateTLSConfig(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on port %v\n", port)

	for {
		connection, err := listener.Accept(context.Background())

		if err != nil {
			log.Fatal(err)
		}
		go handler(connection)
	}
}

// TODO: move eror functions somewhere
func (c *Connections) handlePublisher(connection quic.Connection) {
	log.Printf("New publisher connected: %v\n", connection.RemoteAddr())
	c.addPublisher(connection)
	log.Println(len(c.publishers), "pub length")
	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		go func(stream quic.Stream) {
			defer func() {
				log.Printf("Publisher stream closed: %v\n", stream.StreamID())
				c.removePublisher(connection)
			}()
			buf := make([]byte, 1024)
			for {
				n, err := stream.Read(buf)
				log.Println(string(buf[:n]))
				if err != nil {
					log.Println(err)
					return
				}
				//s.broadcastMessage(buf[:n])
			}
		}(stream)
	}
}

func (c *Connections) handleSubscriber(connection quic.Connection) {
	log.Printf("New subscriber connected: %v\n", connection.RemoteAddr())
	c.addSubscriber(connection)
	log.Println(len(c.subscribers), "sub length")
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

// TODO: move to different file
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"h3-23"},
	}
}
