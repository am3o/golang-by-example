package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/minio/highwayhash"
)

func hashMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return base64.RawURLEncoding.EncodeToString(hash[:16])
}

func hashHighway64(input string) string {
	key, err := hex.DecodeString("totally unknown") // use your own key here
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return ""
	}

	hash, err := highwayhash.New64(key)
	if err != nil {
		fmt.Printf("Failed to create HighwayHash instance: %v", err) // add error handling
		return ""
	}

	if _, err = io.Copy(hash, bytes.NewReader([]byte(input))); err != nil {
		fmt.Printf("Failed to read from file: %v", err) // add error handling
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func hashHighway256(input string) string {
	key, err := hex.DecodeString("totally unknown") // use your own key here
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return ""
	}

	hash, err := highwayhash.New(key)
	if err != nil {
		fmt.Printf("Failed to create HighwayHash instance: %v", err) // add error handling
		return ""
	}

	if _, err = io.Copy(hash, bytes.NewReader([]byte(input))); err != nil {
		fmt.Printf("Failed to read from file: %v", err) // add error handling
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	fmt.Println(hashHighway64("abcdef"))
}
