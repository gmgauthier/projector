package main

import (
	"os"
	"path/filepath"
)

func createDir(fpath string) error {
	if err := os.Mkdir(filepath.Dir(fpath + string(filepath.Separator)), 0770); err != nil {
		return err
	}
	return nil
}

// for generating the project blanks
func createFile(fpath string) error {
	// create directory automatically, if non-existent
	if err := os.MkdirAll(filepath.Dir(fpath), 0770); err != nil {
		return err
	}
	//create the file, but that's all
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
