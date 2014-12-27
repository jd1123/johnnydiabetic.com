package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: hascheck <password>")
	}
	str := os.Args[1]
	hpass, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("Base64'd hash:", base64.StdEncoding.EncodeToString(hpass))
}
