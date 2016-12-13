package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	csvfile      string = "DSMM_CSV_input_file/dsmm_assessments.csv"
	singleRecord DsmmAssessmentRecord
	allRecords   []DsmmAssessmentRecord
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

// Create file directories
func prepDirs() {
	os.Mkdir("./output", 0777)
}

func main() {
	log.Printf("Working digging up files...")
	t0 := time.Now()

	allRecords := getCsvData(csvfile)
	prepDirs()

	// Loop over all CSV records and process one by one
	for _, singleRecord := range allRecords {
		writeCsvDataToWordDoc(singleRecord)
		parseCsvDataToPptxXml(singleRecord)
	}

	t1 := time.Now()
	log.Printf("The program took %v minutes to run.\n", t1.Sub(t0).Minutes())

} // end main()
