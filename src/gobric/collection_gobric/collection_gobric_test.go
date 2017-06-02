package main

import (
	"testing"
)

var testFile string = "./testfiles/woa13_95A4_s00_01_bad.xml"

func TestReadMrrCsvSourceFile(t *testing.T) {
	testData := readMrrCsvSourceFile()
	if testData == nil {
		t.Error("No data, got ", testData)
	}
}

func TestGranuleRulesRunRubric(t *testing.T) {
	xpathFound, xpathNotFound := 13.00, 2.00
	actual1, actual2 := granuleRulesRunRubric(testFile)
	if xpathFound != actual1 && xpathNotFound != actual2 {
		t.Error("Expected 13, 2 as return values, got %d and %d", actual1, actual2)
	}
}

func TestCalculateRubricScore(t *testing.T) {
	test1, test2 := 15.00, 0.00
	actual := calculateRubricScore(test1, test2)
	if actual != 100.00 {
		t.Error("Expected 80 as return value, got ", actual)
	}
	test1, test2 = 4.00, 1.00
	actual = calculateRubricScore(test1, test2)
	if actual != 80.00 {
		t.Error("Expected 80 as return value, got ", actual)
	}
	test1, test2 = 87.00, 13.00
	actual = calculateRubricScore(test1, test2)
	if actual != 87.00 {
		t.Error("Expected 87 as return value, got ", actual)
	}
	test1, test2 = 877.00, 33.00
	actual = calculateRubricScore(test1, test2)
	if actual != 96.37362637362638 {
		t.Error("Expected 96.37362637362638 as return value, got ", actual)
	}
	// test with wrong outcome 86
	test1, test2 = 87.00, 13.00
	actual = calculateRubricScore(test1, test2)
	if actual == 86.00 {
		t.Error("Expected 87 as return value, got ", actual)
	}
}
