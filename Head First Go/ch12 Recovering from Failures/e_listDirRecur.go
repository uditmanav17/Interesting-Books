package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDir(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		filepath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDir(filepath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filepath)
		}
	}
	return nil
}

func main() {
	dir := "../ch11 Interfaces"
	err := scanDir(dir)
	if err != nil {
		log.Fatal(err)
	}
}
