package main

import (
	"flag"
	"log"
	"ssl-check/internal/sslcheck"
)

func main() {
	serverName := flag.String("server", "", "The server name to check the SSL certificate for (required).")
	flag.Parse()

	if *serverName == "" {
		flag.Usage()
		log.Fatal("Error: The -server flag is required.")
	}

	err := sslcheck.CheckCertificate(*serverName)
	if err != nil {
		log.Fatalf("Error checking certificate: %v", err)
	}
}
