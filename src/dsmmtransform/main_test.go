package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	testFile string = "DSMM_CSV_input_file/testfile/dsmm_assessments.csv"
)

func init() {
	os.Mkdir("./output", 0777)
}

func TestGetCsvData(t *testing.T) {
	allRecords = getCsvData(testFile)
	if len(allRecords) != 764 {
		t.Error("Expected CSV data slice to be 764, got ", len(allRecords))
	}
}

func TestWriteCsvDataToWordDoc(t *testing.T) {
	for _, singleRecord := range allRecords[1:4] {
		writeCsvDataToWordDoc(singleRecord)
	}
	testData, _ := ioutil.ReadDir("./output")
	if len(testData) != 3 {
		t.Error("Expected 3 files in output dir, got ", len(testData))
	}
}

func TestParseCsvDataToPptxXml(t *testing.T) {
	for _, singleRecord := range allRecords[1:4] {
		parseCsvDataToPptxXml(singleRecord)
		if _, err := os.Stat("./output/" + singleRecord.C + "_Star_rating_template.pptx"); os.IsNotExist(err) {
			t.Error("expected %s_Star_rating_template.pptx in output directory", singleRecord.C)
		}
	}
}
