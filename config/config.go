package config

import (
	"encoding/base64"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Various config used throughout the webapp
// FIXME: Is it safe to store the keys in memory???
var CONFIG = map[string]string{
	"homeDir":           HomeDir(),
	"templateDir":       AddDir("templates"),
	"staticDir":         AddDir("static"),
	"dbAddress":         "localhost",
	"dbLogin":           "",
	"dbPW":              "",
	"authenticationKey": "1b9H1muvPLWByNWvYSKg7RQ6eEqpeUoAZMomID/Emng=", // base64
	"encryptionKey":     "gT3iJG5+WpTqgmmQnHB1XZfX3PTfWVe68l0iEhZCUg0=", // base64
	"dbName":            "test",
	"usersCollection":   "users",
	"blogCollection":    "blogposts",
}

func HomeDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}

func AddDir(p string) string {
	h := HomeDir()
	return path.Join(h, p)
}

func AuthenticationKey() []byte {
	key, err := base64.StdEncoding.DecodeString(CONFIG["authenticationKey"])
	if err != nil {
		log.Println("Encoding Error")
		return nil
	}
	return key
}

func EncryptionKey() []byte {
	key, err := base64.StdEncoding.DecodeString(CONFIG["authenticationKey"])
	if err != nil {
		log.Println("Encoding Error")
		return nil
	}
	return key
}
