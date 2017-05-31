package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	//"io"
	//"log"
	"os"
	//"reflect"
	//"strings"
)

var (
	mrr_source string = "/nodc/users/tjaensch/belay/belay/fixtures/mrr.csv"
	data *csv.Reader = readMrrCsvSourceFile()
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

func readMrrCsvSourceFile() *csv.Reader {
	mrr, err := os.Open(mrr_source)
	checkError("Open mrr.csv failed", err)
	return csv.NewReader(bufio.NewReader(mrr))
	/* for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
	} */
}
