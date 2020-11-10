package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fatal(err error) {
	// if err != nil {
	// 	panic(errors.Wrap(err, err.Error()))
	// }
	if err != nil {
		// handle error
		fmt.Errorf("Error: %v", err)
		os.Exit(1)
	}
}

type lWriter struct{}

func (lWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("wrote %v bytes", len(bs))
	fmt.Println()
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://example.com/")
	Fatal(err)
	// io.Copy(os.Stdout, resp.Body)
	w := lWriter{}
	io.Copy(w, resp.Body)

	// reader := make([]byte, 99999)
	// resp.Body.Read(reader)
	// fmt.Println(string(reader))
	fmt.Println(resp.StatusCode)
}
