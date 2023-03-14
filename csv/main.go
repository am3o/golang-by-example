package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Sensor struct {
	ID      int
	Name    string
	Type    string
	Reading float64
	Units   string
	Event   string
	Status  float64
}

func ReadFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", file), nil
}

func main() {
	raw, err := ReadFile("raw.csv")
	if err != nil {
		os.Exit(2)
	}

	sensors := ParseCSV(raw)

	for _, sensor := range sensors {
		fmt.Println(sensor)
	}
}

func ParseCSV(content string) (sensors []Sensor) {
	reader := csv.NewReader(strings.NewReader(content))
	reader.Comma = '|'
	reader.FieldsPerRecord = 6

	for {
		record, err := reader.Read()
		if err != nil && err == io.EOF {
			break
		}

		if len(record) >= 6 &&
			strings.TrimSpace(record[3]) != "N/A" &&
			strings.TrimSpace(record[1]) != "Name" {
			sensors = append(sensors, Sensor{
				ID: func() int {
					id, err := strconv.Atoi(strings.TrimSpace(record[0]))
					if err != nil {
						panic(err)
					}
					return id
				}(),
				Name: record[1],
				Type: record[2],
				Reading: func() float64 {
					reading, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
					if err != nil {
						panic(err)
					}
					return reading
				}(),
				Units: record[4],
				Event: record[5],
			})
		}
	}
	return sensors
}
