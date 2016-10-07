package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFile string = "./nc/xerb1_197002.nc"

func init() {
	os.Mkdir("./ncml", 0777)
}

func TestGetFileName(t *testing.T) {
	expected := "xerb1_197002"
	actual := getFileName(testFile)
	if expected != actual {
		t.Error("Expected 'xerb1_197002', got ", actual)
	}
}

func TestFindNcFiles(t *testing.T) {
	result := findNcFiles("./nc")
	if len(result) <= 0 {
		t.Error("Got ", result)
	}
}

func TestNcdump(t *testing.T) {
	ncdump(testFile)
	file, err := os.Stat("./ncml/" + strings.TrimSuffix(filepath.Base(testFile), ".nc") + ".ncml")
	if err != nil {
		t.Error("File not found, got", file)
	}
}
