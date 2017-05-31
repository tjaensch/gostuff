package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	//"log"
	"os"
	"reflect"
	//"strings"
)

var (
	mrr_source string = "/nodc/users/tjaensch/belay/belay/fixtures/mrr.csv"
	data io.Reader = openCsvSourceFile()
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

func openCsvSourceFile() io.Reader {
	mrr, err := os.Open(mrr_source)
	checkError("Open mrr.csv failed", err)
	return mrr
}

func readMrrCsvSourceFile() {
	r := csv.NewReader(bufio.NewReader(data))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(reflect.TypeOf(record[15]))
	}
}
