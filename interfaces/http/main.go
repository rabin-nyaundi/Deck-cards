package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("http request")

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
