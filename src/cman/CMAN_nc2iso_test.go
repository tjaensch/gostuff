package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFile string = "/nodc/web/data.nodc/htdocs/ndbc/cmanwx/2016/05/NDBC_44020_201605_D6_v00.nc"

func init() {
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./xml_output", 0777)
	os.Mkdir("./netcdf3", 0777)
}

func TestGetFileName(t *testing.T) {
	expected := "NDBC_44020_201605_D6_v00"
	actual := getFileName(testFile)
	if expected != actual {
		t.Error("Expected 'NDBC_44020_201605_D6_v00', got ", actual)
	}
}

func TestGetEnglishTitle(t *testing.T) {
	expected := "NDBC-CMANWx_44020_201605_D6_v00 - C-MAN/Wx buoy 44020 for 201605, deployment 6"
	actual := getEnglishTitle(testFile)
	if expected != actual {
		t.Error("Expected 'NDBC-CMANWx_44020_201605_D6_v00 - C-MAN/Wx buoy 44020 for 201605, deployment 6', got ", actual)
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
	expected := "cmanwx/2016/05/"
	actual := getFilePath(testFile)
	if strings.Contains(actual, expected) != true {
		t.Fatalf("File paths don't match %s: %s", expected, actual)
	}
}

func TestFindNcFiles(t *testing.T) {
	result := findNcFiles("/nodc/web/data.nodc/htdocs/ndbc/cmanwx/")
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
	additions := ncmlAdditions{ncFileName: getFileName(testFile), fileSize: getFileSize(testFile), dataPath: getFilePath(testFile), englishTitle: getEnglishTitle(testFile)}
	appendToNcml(testFile, additions)
	input, _ := ioutil.ReadFile("./ncml/" + getFileName(testFile) + ".ncml")
	if !strings.Contains(string(input), "/ndbc/cmanwx/2016/05/") {
		t.Error("appendtoNcml tanked")
	}
}

func TestXsltprocToISO(t *testing.T) {
	xsltprocToISO(testFile, "/nodc/users/tjaensch/xsl/cman/XSL/ncml2iso_modified_from_UnidataDD2MI_CMAN_Thomas_edits.xsl")
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
	if !strings.Contains(string(input), "gov.noaa.nodc:NDBC-CMANWx") {
		t.Error("addCollectionMetadata tanked")
	}
}
