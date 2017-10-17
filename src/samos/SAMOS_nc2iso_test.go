package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFile string = "/nodc/web/data.nodc/htdocs/coaps/samos/WTEJ/2010/08/WTEJ_20100808v30001.nc"

var testFiles []string = []string{
	"/nodc/web/data.nodc/htdocs/coaps/samos/KAOU/2011/06/KAOU_20110613v10001.nc",
	"/nodc/web/data.nodc/htdocs/coaps/samos/WBP3210/2011/10/WBP3210_20111008v10001.nc",
	"/nodc/web/data.nodc/htdocs/coaps/samos/WTDF/2010/03/WTDF_20100318v10001.nc",
	"/nodc/web/data.nodc/htdocs/coaps/samos/WTEK/2012/07/WTEK_20120714v20001.nc",
	"/nodc/web/data.nodc/htdocs/coaps/samos/ZCYL5/2016/06/ZCYL5_20160630v30001.nc",
}

func init() {
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./xml_output", 0777)
}

func TestNc2iso(t *testing.T) {
	nc2iso(testFiles)
	files, _ := ioutil.ReadDir("./xml_output")
	if len(files) != 5 {
		t.Error("Expected 5 files, got ", len(files))
	}
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
	additions := ncmlAdditions{ncFileName: getFileName(testFile), fileSize: getFileSize(testFile), dataPath: getFilePath(testFile)}
	appendToNcml(testFile, additions)
	input, _ := ioutil.ReadFile("./ncml/" + getFileName(testFile) + ".ncml")
	if !strings.Contains(string(input), getFilePath(testFile)) {
		t.Error("appendtoNcml tanked")
	}
}

func TestXsltprocToISO(t *testing.T) {
	xsltprocToISO(testFile, "/nodc/users/tjaensch/xsl.git/samos/XSL/ncml2iso_SAMOS_Thomas_edits.xsl")
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
