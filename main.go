package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	filename = "data.csv"
)

// CsvLine holds csv parsed data
type CsvLine struct {
	Column1 string
	Column2 string
}

func readCSVFile() []CsvLine {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read file into variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	data := []CsvLine{}
	for _, line := range lines {
		data = append(data, CsvLine{
			Column1: line[0],
			Column2: line[1],
		})
	}
	return data
}

func main() {
	data := readCSVFile()
	fmt.Println(data)
}
