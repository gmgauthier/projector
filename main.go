package main

import (
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	dirlist, _ := os.ReadDir(wd)
	fmt.Println(wd)
	for _, item := range dirlist {
		if os.DirEntry.IsDir(item) {
			fmt.Println(wd + "/" + os.DirEntry.Name(item) + "/")
		} else {
			fmt.Println(wd + "/" + os.DirEntry.Name(item))
		}
	}
}
