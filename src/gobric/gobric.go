package main

import (
	"bufio"
	//"bytes"
	"encoding/csv"
	"fmt"
	"io"
	//"log"
	"os"
	"os/exec"
	//"reflect"
	"strings"
)

var (
	mrr_source string      = "/nodc/users/tjaensch/belay/belay/fixtures/mrr.csv"
	data       *csv.Reader = readMrrCsvSourceFile()
	xmlFile    string      = "./testfiles/woa13_95A4_s00_01.xml"
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func main() {
	granuleRulesRunRubric(xmlFile)
}

func readMrrCsvSourceFile() *csv.Reader {
	mrr, err := os.Open(mrr_source)
	checkError("Open mrr.csv failed", err)
	return csv.NewReader(bufio.NewReader(mrr))
}

func granuleRulesRunRubric(xmlFile string) {
	var xmlNode []byte
	for {
		column, err := data.Read()
		if err == io.EOF {
			break
		}
		// line[6] contains the granule "Not/Required" information in CSV data file
		if !strings.Contains(column[6], "Not") && strings.Contains(column[6], "Required") {
			cmdName := "xpath"
			// column[16] contains the xpath expressions
			cmdArgs := []string{xmlFile, column[16]}
			if xmlNode, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
				fmt.Printf(xmlFile+": Something went wrong with running xpath, program exiting.", err)
				os.Exit(1)
			}
			if string(xmlNode) == "" {
				fmt.Println("Not found: ", column[16])
			} else {
				fmt.Println("Found: " + column[16] + " --> " + string(xmlNode))
			}
		} else {
			continue
		}
	}
}
