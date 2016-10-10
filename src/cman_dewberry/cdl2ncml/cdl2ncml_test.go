package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFile string = "/nodc/projects/buoy/dewberry_cdl/Allf291Files/197002/xerb1_197002.cdl"

func init() {
	os.Mkdir("./nc", 0777)
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./cdl2ncml_conversion_failures", 0777)
}

func TestGetFileName(t *testing.T) {
	expected := "xerb1_197002"
	actual := getFileName(testFile)
	if expected != actual {
		t.Error("Expected 'xerb1_197002', got ", actual)
	}
}

func TestFindCdlFiles(t *testing.T) {
	result := findCdlFiles("/nodc/projects/buoy/dewberry_cdl/Allf291Files/")
	if len(result) <= 0 {
		t.Error("Got ", result)
	}
}

func TestNcgen(t *testing.T) {
	ncgen(testFile)
	file, err := os.Stat("./nc/" + strings.TrimSuffix(filepath.Base(testFile), ".cdl") + ".nc")
	if err != nil {
		t.Error("File not found, got", file)
	}
}

func TestNcdump(t *testing.T) {
	ncdump(testFile)
	file, err := os.Stat("./ncml/" + strings.TrimSuffix(filepath.Base(testFile), ".cdl") + ".ncml")
	if err != nil {
		t.Error("File not found, got", file)
	}
}
