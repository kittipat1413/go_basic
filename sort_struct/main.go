package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Item struct {
	Name  string
	Units string
	Price string
}

func main() {

	lines, err := ReadCsv("data.csv")
	if err != nil {
		panic(err)
	}

	var items []Item
	// Loop through lines & turn into object
	for _, line := range lines {
		item := Item{
			Name:  line[0],
			Units: line[1],
			Price: line[2],
		}
		items = append(items, item)
	}
	fmt.Printf("%v\n", items)
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
