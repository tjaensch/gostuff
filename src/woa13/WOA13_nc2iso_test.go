package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFile string = "/nodc/web/data.nodc/htdocs/nodc/archive/data/0114815/public/salinity/netcdf/95A4/1.00/woa13_95A4_s13_01.nc"

func init() {
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./xml_output", 0777)
}

func TestGetFileName(t *testing.T) {
	expected := "woa13_95A4_s13_01"
	actual := getFileName(testFile)
	if expected != actual {
		t.Error("Expected 'woa13_95A4_s13_01', got ", actual)
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
	expected := "nodc/archive/data/0114815/public/salinity/netcdf/95A4/1.00/"
	actual := getFilePath(testFile)
	if strings.Contains(actual, expected) != true {
		t.Fatalf("File paths don't match %s: %s", expected, actual)
	}
}

func TestFindNcFiles(t *testing.T) {
	result := findNcFiles("/nodc/web/data.nodc/htdocs/nodc/archive/data/0114815/public")
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
	additions := ncmlAdditions{ncFileName: "test", fileSize: 100, datapath: "test/moo/", browsegraphic: "test"}
	appendToNcml(testFile, additions)
	input, _ := ioutil.ReadFile("./ncml/" + getFileName(testFile) + ".ncml")
	if !strings.Contains(string(input), "test/moo/") {
		t.Error("appendtoNcml tanked")
	}
}

func TestXsltprocToISO(t *testing.T) {
	xsltprocToISO(testFile, "/nodc/users/tjaensch/onestop.git/xsl/woa13/XSL/ncml2iso_modified_from_UnidataDD2MI_demo_WOA_Thomas_edits.xsl")
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
	if !strings.Contains(string(input), "doi:10.7289/V5F769GT") {
		t.Error("addCollectionMetadata tanked")
	}
}

func TestGetBrowseGraphicLink(t *testing.T) {
	link := getBrowseGraphicLink(testFile)
	if !strings.Contains(link, "woa13-salinity-seasonal-1degree") {
		t.Error("Got ", link)
	}
}
