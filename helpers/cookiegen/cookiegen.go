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

func GenerateKey(bytes int) string {
	b := make([]byte, bytes)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Println("Io error")
		os.Exit(1)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func main() {

	fmt.Println("128 bit (16 byte) keys:")
	for i := 0; i < 5; i++ {
		fmt.Println("Key", i+1, ":", GenerateKey(16))
	}
	fmt.Println()

	fmt.Println("256 bit (32 byte) keys:")
	for i := 0; i < 5; i++ {
		fmt.Println("Key", i+1, ":", GenerateKey(32))
	}
	fmt.Println()
	fmt.Println("512 bit (64 byte) keys:")
	for i := 0; i < 5; i++ {
		fmt.Println("Key", i+1, ":", GenerateKey(64))
	}
}
