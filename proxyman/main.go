package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	os.Setenv("HTTPS_PROXY", "localhost:9091")
	req, err := http.NewRequest(http.MethodGet, "https://google.de", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Foo", "Bar")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
