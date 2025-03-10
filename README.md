# go-cert-check
Basic check for cert expiration using Go

Go File Structure
```
|-- README.md
|-- cmd
|   `-- ssl-check
|       `-- main.go
|-- go.mod
`-- internal
    `-- sslcheck
        |-- certificate.go
        `-- connection.go
```

Example Execution
``` 
~/source/repos/go-cert-check (main)
$ go run cmd/ssl-check/main.go -server ultimateqa.com
```

Example Output
```
Subject: CN=ultimateqa.com
Issuer: CN=WE1,O=Google Trust Services,C=US
Expiry: Thu, 29 May 2025 07:07:01 UTC
---------------------------
Certificate is valid for 79 more days.
```

Note: repo includes a `launch.config` file to step through `Go` code using Visual Code.