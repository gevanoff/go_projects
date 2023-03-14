package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1] // first argument should be the file name
	fmt.Println("File name: ", filename)
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error. Filename: ", filename, " produced error: ", err)
		os.Exit(1)
	}

	f, err2 := os.Open(filename)

	if err2 != nil {
		fmt.Println("Error. Filename: ", filename, " produced error: ", err2)
		os.Exit(1)
	}
	/*
		bs := make([]byte, 99999)
		f.Read(bs)
		fmt.Println(string(bs))
	*/
	io.Copy(os.Stdout, f) // alternate method
}
