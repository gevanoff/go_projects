package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil { // If error is not nil
		fmt.Println("Error:", err) // print the error
		os.Exit(1)                 // and exit with exit code 1
	}
	// Longish way of reading the response Body data into a byte slice of defined length
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs)) // Render the byte slice as a string and print.
	lw := logWriter{}
	io.Copy(lw, resp.Body) // Use custom type, implementing custom Write function
	//	io.Copy(os.Stdout, resp.Body) // The standard, short way of writing Body data to Stdout
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
