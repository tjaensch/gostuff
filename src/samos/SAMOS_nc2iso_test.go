package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFile string = "/nodc/web/data.nodc/htdocs/coaps/samos/WTEJ/2010/08/WTEJ_20100808v30001.nc"

func init() {
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./xml_output", 0777)
}

func TestGetFileName(t *testing.T) {
	expected := "WTEJ_20100808v30001"
	actual := getFileName(testFile)
	if expected != actual {
		t.Error("Expected 'WTEJ_20100808v30001', got ", actual)
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
	expected := "coaps/samos/WTEJ/2010/08/"
	actual := getFilePath(testFile)
	if strings.Contains(actual, expected) != true {
		t.Fatalf("File paths don't match %s: %s", expected, actual)
	}
}

func TestFindNcFiles(t *testing.T) {
	result := findNcFiles("/nodc/web/data.nodc/htdocs/coaps/samos/")
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
	additions := ncmlAdditions{ncFileName: "test", fileSize: 100, dataPath: "test/moo/blah"}
	appendToNcml(testFile, additions)
	input, _ := ioutil.ReadFile("./ncml/" + getFileName(testFile) + ".ncml")
	if !strings.Contains(string(input), "test/moo/blah") {
		t.Error("appendtoNcml tanked")
	}
}

func TestXsltprocToISO(t *testing.T) {
	xsltprocToISO(testFile, "XSL/ncml2iso_SAMOS_Thomas_edits.xsl")
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
	if !strings.Contains(string(input), "doi:10.7289/V5QJ7F8R") {
		t.Error("addCollectionMetadata tanked")
	}
}
