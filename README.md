# go-cert-check
Basic check for cert expiration using Go

```
$ go run cert-check.go -server=ultimateqa.com
Subject: CN=ultimateqa.com
Issuer: CN=WE1,O=Google Trust Services,C=US
Expiry: Thu, 29 May 2025 07:07:01 UTC
---------------------------
Certificate is valid for 79 more days.
Subject: CN=WE1,O=Google Trust Services,C=US
Issuer: CN=GTS Root R4,O=Google Trust Services LLC,C=US
Expiry: Tue, 20 Feb 2029 14:00:00 UTC
---------------------------
Certificate is valid for 1442 more days.
Subject: CN=GTS Root R4,O=Google Trust Services LLC,C=US
Issuer: CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Expiry: Fri, 28 Jan 2028 00:00:42 UTC
---------------------------
Certificate is valid for 1053 more days.
```