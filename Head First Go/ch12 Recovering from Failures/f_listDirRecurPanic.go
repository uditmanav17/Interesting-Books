package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDir(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filepath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDir(filepath)

		} else {
			fmt.Println(filepath)
		}
	}
	return nil
}

func main() {
	dir := "../ch11 Interfaces"
	scanDir(dir)
}
