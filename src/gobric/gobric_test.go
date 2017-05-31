package main

import (
	"testing"
)

func TestReadMrrCsvSourceFile(t *testing.T) {
	testData := readMrrCsvSourceFile()
	if testData == nil {
		t.Error("No data, got ", testData)
	}
}

func TestGranuleRulesRunRubric(t *testing.T) {
	
}
