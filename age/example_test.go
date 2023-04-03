package main

import (
	"encoding/json"
	"os"
)

type Content struct {
	Id string `json:"ID"`
}

func ExampleContainer_Encryption() {
	file, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	
	passphrase := "secret"

	container, err := NewContainer(file, passphrase)
	if err != nil {
		panic(err)
	}

	writer, err := container.Writer()
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(writer).Encode(Content{Id: "Example"}); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}

	reader, err := container.Reader()
	if err != nil {
		panic(err)
	}

	var content Content
	if err := json.NewDecoder(reader).Decode(&content); err != nil {
		panic(err)
	}
}
