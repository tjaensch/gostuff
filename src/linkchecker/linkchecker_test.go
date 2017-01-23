package main

import (
	"os"
	"testing"
)

func TestFindXmlFiles(t *testing.T) {
	expected := findXmlFiles()
	if len(expected) <= 0 {
		t.Error("Got ", len(expected))
	}
}

func TestFindBrokenHttpLinks(t *testing.T) {
	findBrokenHttpLinks("testfile.xml")
	if _, err := os.Stat("linkchecker_bad_links_log.log"); os.IsNotExist(err) {
		t.Error("testfile.xml should have created log file")
	}
}
