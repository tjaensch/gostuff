package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFile string = "/nodc/web/data.nodc/htdocs/ndbc/co-ops/2014/01/NOS_1612480_201401_D1_v00.nc"

func init() {
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./xml_output", 0777)
}

func TestGetFileName(t *testing.T) {
	expected := "NOS_1612480_201401_D1_v00"
	actual := getFileName(testFile)
	if expected != actual {
		t.Error("Expected 'NOS_1612480_201401_D1_v00', got ", actual)
	}
}

func TestGetEnglishTitle(t *testing.T) {
	expected := "NDBC-COOPS_1612480_201401_D1_v00 - CO-OPS buoy 1612480 for 201401, deployment 1"
	actual := getEnglishTitle(testFile)
	if expected != actual {
		t.Error("Expected 'NDBC-COOPS_1612480_201401_D1_v00 - CO-OPS buoy 1612480 for 201401, deployment 1', got ", actual)
	}
}

func TestGetFileSize(t *testing.T) {
	expected := 1
	actual := getFileSize(testFile)
	if actual < expected {
		t.Error("Expected filesize > 1, got ", actual)
	}
}

func TestGetFilePath(t *testing.T) {
	expected := "co-ops/2014/01/"
	actual := getFilePath(testFile)
	if strings.Contains(actual, expected) != true {
		t.Fatalf("File paths don't match %s: %s", expected, actual)
	}
}

func TestFindNcFiles(t *testing.T) {
	result := findNcFiles("/nodc/web/data.nodc/htdocs/ndbc/co-ops/")
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

func TestAppendToNcml(t *testing.T) {
	additions := ncmlAdditions{ncFileName: "NOS_1612480_201401_D1_v00", fileSize: 100, dataPath: "test/moo/blah", englishTitle: "English Title"}
	appendToNcml(testFile, additions)
	input, _ := ioutil.ReadFile("./ncml/" + getFileName(testFile) + ".ncml")
	if !strings.Contains(string(input), "test/moo/blah") {
		t.Error("appendtoNcml tanked")
	}
}

func TestXsltprocToISO(t *testing.T) {
	xsltprocToISO(testFile, "XSL/ncml2iso_modified_from_UnidataDD2MI_COOPS_Thomas_edits.xsl")
	input, _ := ioutil.ReadFile("./xml_output/" + getFileName(testFile) + ".xml")
	if !strings.Contains(string(input), "<gmi:MI_Metadata") {
		t.Error("xsltprocToISO tanked")
	}
}

func TestAddCollectionMetadata(t *testing.T) {
	addCollectionMetadata(testFile)
	input, err := ioutil.ReadFile("./xml_output/" + getFileName(testFile) + ".xml")
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(input), "gov.noaa.nodc:NDBC-COOPS") {
		t.Error("addCollectionMetadata tanked")
	}
}
