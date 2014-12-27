package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
)

// Simple package to generate random 256 bit keys for gorilla sessions
// Output in a base64 encoded string
func main() {
	authenticationKey := make([]byte, 64)
	encryptionKey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, authenticationKey); err != nil {
		log.Println("Io error")
		os.Exit(1)
	}
	if _, err := io.ReadFull(rand.Reader, encryptionKey); err != nil {
		log.Println("Io error")
		os.Exit(1)
	}
	base64.StdEncoding.EncodeToString(authenticationKey)
	fmt.Println("Authentication key:", base64.StdEncoding.EncodeToString(authenticationKey))
	fmt.Println("Encryption key:", base64.StdEncoding.EncodeToString(encryptionKey))
}
