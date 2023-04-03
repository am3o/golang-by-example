package main

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainer(t *testing.T) {
	const passphrase = "bar"
	expectedContent := []byte("foo\n")

	var file = &bytes.Buffer{}

	container, err := NewContainer(file, passphrase)
	if err != nil {
		t.Fatal(err)
	}

	writer, err := container.Writer()
	assert.NoError(t, err)

	err = writer.Close()
	assert.NoError(t, err)

	reader, err := container.Reader()
	assert.NoError(t, err)

	content, err := io.ReadAll(reader)
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, content)
}

func TestContainerCrypto(t *testing.T) {
	type Content struct {
		Id string `json:"ID"`
	}

	var (
		expectedContent = Content{Id: "A11y"}
	)

	var file = &bytes.Buffer{}
	passphrase := "bar"

	container, err := NewContainer(file, passphrase)
	if err != nil {
		t.Fatal(err)
	}

	writer, err := container.Writer()
	assert.NoError(t, err)

	err = json.NewEncoder(writer).Encode(expectedContent)
	assert.NoError(t, err)

	err = writer.Close()
	assert.NoError(t, err)

	reader, err := container.Reader()
	assert.NoError(t, err)

	var content Content
	err = json.NewDecoder(reader).Decode(&content)
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, content)
}
