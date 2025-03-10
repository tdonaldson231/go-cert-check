package sslcheck

import (
	"fmt"
	"time"
)

func CheckCertificate(serverName string) error {
	_, cert, err := ConnectToServer(serverName)
	if err != nil {
		return err
	}

	if cert != nil {
		fmt.Printf("Subject: %s\n", cert.Subject)
		fmt.Printf("Issuer: %s\n", cert.Issuer)
		fmt.Printf("Expiry: %s\n", cert.NotAfter.Format(time.RFC1123))
		fmt.Println("---------------------------")

		// Check if expiration date is less than 30 days
		durationUntilExpiration := time.Until(cert.NotAfter)
		daysRemaining := int(durationUntilExpiration.Hours() / 24)
		if daysRemaining < 30 {
			fmt.Printf("Certificate will expire in %d days, which is less than 30 days. WARNING!\n", daysRemaining)
		} else {
			fmt.Printf("Certificate is valid for %d more days.\n", daysRemaining)
		}
	} else {
		return fmt.Errorf("no certificate found")
	}

	return nil
}
