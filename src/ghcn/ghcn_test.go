package main

import (
	//"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var testfile string = "AGE00147710"

func TestDownloadStationsTextFile(t *testing.T) {
	downloadStationsTextFile()
	file, err := os.Stat("ghcnd-stations.txt")
	if err != nil {
		t.Error("File not found, got", file)
	}
}

func TestReadInStationsFileInfo(t *testing.T) {
	_, latMap, lonMap, stationLongNameMap := readInStationsFileInfo()
	if len(latMap) != len(lonMap) {
		t.Error("len(latMap) doesn't match len(lonMap)")
	}
	if len(stationLongNameMap) < 103000 {
		t.Error("len(stationLongNameMap) not right, got", len(stationLongNameMap))
	}
}

func TestGetIndividualDataFileAsString(t *testing.T) {
	expected := getIndividualDataFileAsString(testfile)
	if reflect.TypeOf(expected).Kind() != reflect.String {
		t.Error("expected string type, got", reflect.TypeOf(expected))
	}
}

func TestGetMetadataKeywordsForStationFile(t *testing.T) {
	metadataKeywordsForStationFile := getMetadataKeywordsForStationFile(testfile)
	if len(metadataKeywordsForStationFile) < 1 {
		t.Error("len(metadataKeywordsForStationFile) not right, got", len(metadataKeywordsForStationFile))
	}
}
