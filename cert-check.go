package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	serverName := flag.String("server", "", "The server name to check the SSL certificate for (required).")
	port := "443"
	flag.Parse()

	// if serverName is empty, display usage and exit
	if *serverName == "" {
		flag.Usage()
		log.Fatal("Error: The -server flag is required.")
	}

	conn, err := tls.Dial("tcp", *serverName+":"+port, &tls.Config{
		InsecureSkipVerify: false,
	})
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	state := conn.ConnectionState()

	// Iterate through the certificates
	for _, cert := range state.PeerCertificates {
		// Print certificate details
		fmt.Printf("Subject: %s\n", cert.Subject)
		fmt.Printf("Issuer: %s\n", cert.Issuer)
		fmt.Printf("Expiry: %s\n", cert.NotAfter.Format(time.RFC1123))
		fmt.Println("---------------------------")

		// Verify the certificate
		opts := x509.VerifyOptions{
			DNSName: *serverName,
			Roots:   x509.NewCertPool(),
		}
		opts.Roots.AddCert(cert)

		// Check if expiration date is less than 30 days
		daysRemaining := int(cert.NotAfter.Sub(time.Now()).Hours() / 24)
		if daysRemaining < 30 {
			fmt.Printf("Certificate will expire in %d days, which is less than 30 days. WARNING!\n", daysRemaining)
		} else {
			fmt.Printf("Certificate is valid for %d more days.\n", daysRemaining)
		}
	}
}
