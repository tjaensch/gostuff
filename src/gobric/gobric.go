package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	//"log"
	"os"
	//"strings"
)

var (
	mrr_source string = "/nodc/users/tjaensch/belay/belay/fixtures/mrr.csv"
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func main() {
	readMrrCsvSourceFile()
}

func readMrrCsvSourceFile() {
	mrr, err := os.Open(mrr_source)
	checkError("Open mrr.csv failed", err)
	
	data := csv.NewReader(bufio.NewReader(mrr))
	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(len(record[15]))
	}
}
