package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTLSConfigNotNil(t *testing.T) {
	tlsConfig := generateTLSConfig()

	assert.NotNil(t, tlsConfig, "tls.Config has not been generated")
}

func TestGenerateTLSConfigCertificatesCreated(t *testing.T) {
	tlsConfig := generateTLSConfig()

	assert.Equal(t, 1, len(tlsConfig.Certificates), "Number of created tlsConfig.Certificates does not match expected in tls.Config (1)")
}

func TestGenerateTLSConfigNextProtosCreated(t *testing.T) {
	tlsConfig := generateTLSConfig()

	expectedNextProtos := []string{"h3-23"}
	assert.ElementsMatch(t, expectedNextProtos, tlsConfig.NextProtos, "Generated NextProtos does not match expected in the tls.Config")
}
