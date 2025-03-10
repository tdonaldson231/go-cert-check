package sslcheck

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

func ConnectToServer(serverName string) (*tls.Conn, *x509.Certificate, error) {
	conn, err := tls.Dial("tcp", serverName+":443", &tls.Config{
		InsecureSkipVerify: false,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	if len(state.PeerCertificates) > 0 {
		return conn, state.PeerCertificates[0], nil
	}
	return nil, nil, fmt.Errorf("no certificates found")
}
