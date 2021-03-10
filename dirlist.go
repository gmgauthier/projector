package main

import (
	"fmt"
	"os"
)

func dirlist() []string {
	var output []string
	wd, _ := os.Getwd()
	dirlist, _ := os.ReadDir(wd)
	for _, item := range dirlist {
		if os.DirEntry.IsDir(item) {
			output = append(output, fmt.Sprintf(wd + "/" + os.DirEntry.Name(item) + "/"))
		} else {
			output = append(output, fmt.Sprintf(wd + "/" + os.DirEntry.Name(item)))
		}
	}
	return output
}

