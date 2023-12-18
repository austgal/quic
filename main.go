package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/quic-go/quic-go"
)

const pubPort = 6666
const subPort = 6667

func main() {
	go startServer(pubPort)
	go startServer(subPort)

	select {}
}

func startServer(port int) {
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: port})

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
		go handleConnection(connection)
		//TODO: handle pub and sub separately
	}
}

// TODO: move eror functions somewhere
func handleConnection(connection quic.Connection) {
	stream, err := connection.AcceptStream(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		fmt.Println(string(buf[:n]))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func handlePublisher(connection quic.Connection) {
	log.Printf("New publisher connected: %v\n", connection.RemoteAddr())
	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		go func(stream quic.Stream) {

			buf := make([]byte, 1024)
			for {
				n, err := stream.Read(buf)
				fmt.Println(string(buf[:n]))
				if err != nil {
					log.Println(err)
					return
				}
				//s.broadcastMessage(buf[:n])
			}
		}(stream)
	}
}

func handleSubscriber(connection quic.Connection) {
	log.Printf("New subscriber connected: %v\n", connection.RemoteAddr())
	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		go func(stream quic.Stream) {
			defer func() {
				log.Printf("Subscriber stream closed: %v\n", stream.StreamID())
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
