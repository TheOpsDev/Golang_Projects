package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type UrlPair struct {
	Url   string
	Value int
}

type UrlPairList []UrlPair

func (p UrlPairList) Len() int           { return len(p) }
func (p UrlPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p UrlPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// function to rank urls based on mapped values
func rankByUrlValue(rowsMap map[string]int) UrlPairList {
	topTen := make(UrlPairList, len(rowsMap))
	i := 0

	for k, v := range rowsMap {
		topTen[i] = UrlPair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(topTen))
	return topTen
}

// function to validate file param
func validateFile(file *os.File) error {
	fStat, err := file.Stat()
	if err != nil {
		return err
	}

	// check if file is empty
	if fStat.Size() == 0 {
		log.Fatal("File is empty")
	}
	return nil
}

// function read file concurrently
func readFile(file *os.File) chan []string {
	rows := make(chan []string)
	go func() {
		parser := csv.NewReader(file)
		parser.Comma = ' '

		defer close(rows)

		for {
			record, err := parser.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			rows <- record
		}
	}()

	return rows
}

// function to convert read rows to map
func processRows(rows chan []string) map[string]int {
	rowsMap := make(map[string]int)

	for row := range rows {
		i, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		rowsMap[row[0]] = i
	}

	return rowsMap
}

func main() {
	var filename string
	var count int

	// read user arguments
	flag.StringVar(&filename, "filename", "", "file path to read from")
	flag.Parse()

	if filename == "" {
		log.Fatal("No filepath was provided. Use -filename flag to pass in value")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// validate file
	err = validateFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// process and sort data
	rows := readFile(file)
	mappedRows := processRows(rows)
	sortedResults := rankByUrlValue(mappedRows)

	if len(sortedResults) < 10 {
		count = len(sortedResults)
	} else {
		count = 10
	}

	// output results
	for i := 0; i < count; i++ {
		fmt.Println(sortedResults[i].Url)
	}
}
