package main

import (
  "strings"
	"testing"
)

var (
	testFile1 string = "./testfiles/test1.txt"
	testFile2 string = "./testfiles/test2.txt"
)

func TestCompareFiles(t *testing.T) {
	compareFiles(testFile1, testFile2)
  if !strings.Contains(readFile("compareFiles_log.log"), "time10 = 4452") {
    t.Fail()
  }
}

func TestReadFile(t *testing.T) {
  expected1 := readFile(testFile1)
  expected2 := readFile(testFile2)
  expected3 := readFile(testFile1)
  if expected1 == expected2 {
		t.Fail()
	}
  if expected1 != expected3 {
    t.Fail()
  }
}
