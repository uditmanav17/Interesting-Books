package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// added reportPanic
func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

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
	defer reportPanic()
	dir := "../ch11 Interfaces"
	panic("some issue")
	scanDir(dir)
}
