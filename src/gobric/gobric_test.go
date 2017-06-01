package main

import (
	"testing"
)

var testFile string = "./testfiles/woa13_95A4_s00_01.xml"

func TestReadMrrCsvSourceFile(t *testing.T) {
	testData := readMrrCsvSourceFile()
	if testData == nil {
		t.Error("No data, got ", testData)
	}
}

func TestGranuleRulesRunRubric(t *testing.T) {
	xpathFound, xpathNotFound := 15, 0
	actual1, actual2 := granuleRulesRunRubric(testFile)
	if xpathFound != actual1 && xpathNotFound != actual2 {
		t.Error("Expected 15, 0 as return values, got %d and %d", actual1, actual2)
	}
}
