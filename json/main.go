package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Content struct {
	T Number `json:"type"`
}

type Number float64

func (d *Number) UnmarshalJSON(v []byte) error {
	value, err := strconv.ParseFloat(strings.Trim(string(v), "\""), 10)
	if err != nil {
		return err
	}

	*d = Number(value + 0.10)
	return nil
}

func main() {
	byt := []byte(`{"type":6.13}`)
	if err := json.NewEncoder(os.Stdout).Encode(byt); err != nil {
		panic(err)
	}
}
