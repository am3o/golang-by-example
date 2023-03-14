package main

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Tbody struct {
	Rows []struct {
		Data []string `xml:"td"`
	} `xml:"tr"`
}

func main() {
	raw, err := os.ReadFile("./raw.html")
	if err != nil {
		panic(err)
	}

	var content Tbody
	if err := xml.NewDecoder(bytes.NewReader(raw)).Decode(&content); err != nil {
		panic(err)
	}

	file, err := os.CreateTemp(os.TempDir(), "*.csv")
	if err != nil {
		panic(err)
	}
	defer fmt.Println(file.Name())

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()

	fmt.Printf("lines to write: %d\n", len(content.Rows))
	for _, value := range content.Rows {
		for idx := range value.Data {
			value.Data[idx] = strings.ReplaceAll(value.Data[idx], "\n", " ")
			value.Data[idx] = strings.ReplaceAll(value.Data[idx], "\t", "")
		}

		if err := writer.Write(value.Data); err != nil {
			panic(err)
		}
	}
}
