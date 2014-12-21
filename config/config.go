package config

import (
	"os"
	"path"
	"path/filepath"
)

var CONFIG = map[string]string{
	"homeDir":     HomeDir(),
	"templateDir": AddDir("templates"),
	"staticDir":   AddDir("static"),
	"db":          "",
	"dbLogin":     "",
	"dbPW":        "",
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
