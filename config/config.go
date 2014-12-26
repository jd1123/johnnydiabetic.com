package config

import (
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
	"db":                "",
	"dbLogin":           "",
	"dbPW":              "",
	"authenticationKey": "", // base64 encoded string
	"encryptionKey":     "", // base64 encoded string
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
