package main

import (
	"testing"
)

func TestGenerateTLSConfigNotNil(t *testing.T) {
	tlsConfig := generateTLSConfig()

	if tlsConfig == nil {
		t.Error("tls.Config has not been generated")
	}
}

func TestGenerateTLSConfigCertificatesCreated(t *testing.T) {
	tlsConfig := generateTLSConfig()

	if len(tlsConfig.Certificates) != 1 {
		t.Error("Number of created tlsConfig.Certificates does not match expected in tls.Config")
	}
}

func TestGenerateTLSConfigNextProtosCreated(t *testing.T) {
	tlsConfig := generateTLSConfig()

	expectedNextProtos := []string{"h3-23"}
	if len(tlsConfig.NextProtos) != len(expectedNextProtos) {
		t.Errorf("Unexpected NextProtos in the tls.Config. Expected: %v, Got: %v", expectedNextProtos, tlsConfig.NextProtos)
	}
}
