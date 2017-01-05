package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	testFile string = "/nodc/PROJECTS/satdata/OER/Metadata/waf/EX1004L2_VID_20100629T021804Z_CPHD_FIRST_BOTTOM.mov.xml"
	)

func init() {
	os.Mkdir("./isolite_xml_output", 0777)
}

func TestGetFileName(t *testing.T) {
	expected := "EX1004L2_VID_20100629T021804Z_CPHD_FIRST_BOTTOM.mov"
	actual := getFileName(testFile)
	if expected != actual {
		t.Error("Expected 'EX1004L2_VID_20100629T021804Z_CPHD_FIRST_BOTTOM.mov', got ", actual)
	}
}

func TestFindXmlFiles(t *testing.T) {
	result := findXmlFiles(xmlFilePath)
	if len(result) <= 0 {
		t.Error("Got ", result)
	}
}

func TestXsltprocToISO(t *testing.T) {
	xsltprocToISO(testFile, xslFile)
	input, _ := ioutil.ReadFile("./isolite_xml_output/" + getFileName(testFile) + ".xml")
	if !strings.Contains(string(input), "OER.EX1004L2_DIVE01_20100629T021804Z_CPHD_FIRST_BOTTOM") {
		t.Error("xsltprocToISO tanked")
	}
}

/* func TestAddCollectionMetadata(t *testing.T) {
	addCollectionMetadata(testFile)
	input, err := ioutil.ReadFile("./isolite_xml_output/" + getFileName(testFile) + ".xml")
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(input), "TBD") {
		t.Error("addCollectionMetadata tanked")
	} 
} */
