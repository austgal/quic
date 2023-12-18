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

const publisherPort = 6666
const subscriberPort = 6667

func main() {
	startServer()
}

func startServer() {
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: publisherPort})

	tr := quic.Transport{
		Conn: udpConn,
	}
	listener, err := tr.Listen(generateTLSConfig(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on port %d\n", publisherPort)

	for {
		connection, err := listener.Accept(context.Background())
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
