package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct{}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fw := fileWriter{}
	io.Copy(fw, file)
}

func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
