package main

import (
	"fmt"
	"os"
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func main() {

	allRecords := getCsvData()
	writeCsvDataToWordDoc(allRecords)

}
