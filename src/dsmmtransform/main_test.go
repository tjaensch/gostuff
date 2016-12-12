package main

import (
  "io/ioutil"
	"os"
	"testing"
)

var (
  testFile string = "DSMM_CSV_input_file/testfile/dsmm_assessments.csv"
  files []DsmmAssessmentRecord
)

func init() {
	os.Mkdir("./output", 0777)
}

func TestGetCsvData(t *testing.T) {
  files = getCsvData(testFile)
   if len(files) != 764 {
    t.Error("Expected CSV data slice to be 764, got ", len(files))
  }
}

func TestWriteCsvDataToWordDoc(t *testing.T) {
  writeCsvDataToWordDoc(files[1:4])
  testData, _ := ioutil.ReadDir("./output")
  if len(testData) != 3 {
   t.Error("Expected 3 files in output dir, got ", len(testData))
 }
}
