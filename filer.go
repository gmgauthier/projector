package main

import (
	"log"
	"os"
	"path/filepath"
)

// for generating the project blanks
func createFile(fpath string) bool {
	fl, err := newFile(fpath)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fl.Close()
	return true
}

// in case we need a file we can edit
func newFile(fpath string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(fpath), 0770); err != nil {
		return nil, err
	}
	return os.Create(fpath)
}
